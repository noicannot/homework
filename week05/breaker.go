package main

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/context"
)

var (
	BREAKER_OPEN_ERROR error = errors.New("breaker_open")
	OPEN_TO_HALF_ERROR error = errors.New("OPEN_TO_HALF")
)

var (
	FUNC_NIL_ERROR error = errors.New("runFunc nil")
	NAME_NIL_ERROR error = errors.New("name nil")
)

type breaker struct {
	name        string
	sleepWindow int64
	cycleTime   int64
	counter     *SlidingWindow
	lpm         *limitPoolManager
}

type breakerManager struct {
	mutex   *sync.RWMutex
	manager map[string]*breaker
}

var bm *breakerManager

func init() {
	bm = new(breakerManager)
	bm.manager = make(map[string]*breaker)
	bm.mutex = &sync.RWMutex{}
}

type runFunc func() error

type fallbackFunc func(error)

func newBreaker(b *breakSettingInfo) *breaker {
	lpm := NewLimitPoolManager(b.BreakerTestMax)
	counter := NewSlidingWindow(SlidingWindowSetting{CycleTime: b.Interval,
		ErrorPercent:      b.ErrorPercentThreshold,
		BreakErrorPercent: b.BreakerErrorPercentThreshold,
		BreakCnt:          b.BreakerTestMax,
	})
	return &breaker{
		name:        b.Name,
		cycleTime:   time.Now().Local().Unix() + b.SleepWindow,
		sleepWindow: b.SleepWindow,
		counter:     counter,
		lpm:         lpm,
	}
}

func getBreakerManager(name string) (*breaker, error) {
	if name == "" {
		return nil, errors.New("no name")
	}
	bm.mutex.RLock()
	pBreaker, ok := bm.manager[name]
	if !ok {
		bm.mutex.RUnlock()
		bm.mutex.Lock()
		defer bm.mutex.Unlock()
		if pBreaker, ok := bm.manager[name]; ok {
			return pBreaker, nil
		}
		pbreak, err := NewBreakSettingInfo().SetName(name).AddBreakSetting()
		if err != nil {
			return nil, err
		}
		bm.manager[name] = pbreak
		return pbreak, nil
	} else {
		defer bm.mutex.RUnlock()
		return pBreaker, nil
	}
}

func (this *breaker) fail() {
	state := this.counter.GetStatus()
	switch state {
	case STATE_CLOSED:
		atomic.StoreInt64(&this.cycleTime, time.Now().Local().Unix()+this.sleepWindow)
		this.counter.Add(false)
	case STATE_OPEN:
		if time.Now().Local().Unix() > atomic.LoadInt64(&this.cycleTime) {
			if this.counter.AddBreak(false) {
				defer this.lpm.ReturnAll()
				atomic.StoreInt64(&this.cycleTime, time.Now().Local().Unix()+this.sleepWindow)
			}
		}
	}
}

func (this *breaker) success() {
	state := this.counter.GetStatus()
	switch state {
	case STATE_CLOSED:
		atomic.StoreInt64(&this.cycleTime, time.Now().Local().Unix()+this.sleepWindow)
		this.counter.Add(true)
	case STATE_OPEN:
		if time.Now().Local().Unix() > atomic.LoadInt64(&this.cycleTime) {
			if this.counter.AddBreak(true) {
				defer this.lpm.ReturnAll()
				atomic.StoreInt64(&this.cycleTime, time.Now().Local().Unix()+this.sleepWindow)
			}
		}
	}
}

func (this *breaker) safelCalllback(fallback fallbackFunc, err error) {
	if fallback == nil {
		return
	}
	fallback(err)
}

func safelCalllback(fallback fallbackFunc, err error) {
	if fallback == nil {
		return
	}
	fallback(err)
}

func (this *breaker) beforeDo(ctx context.Context, name string) error {
	switch this.counter.GetStatus() {
	case STATE_OPEN:
		if this.cycleTime < time.Now().Local().Unix() {
			return OPEN_TO_HALF_ERROR
		}
		return BREAKER_OPEN_ERROR
	}
	return nil
}

func (this *breaker) afterDo(ctx context.Context, run runFunc, fallback fallbackFunc, err error) error {
	switch err {
	case BREAKER_OPEN_ERROR:
		this.safelCalllback(fallback, BREAKER_OPEN_ERROR)
		return nil
	case OPEN_TO_HALF_ERROR:
		if !this.lpm.GetTicket() {
			this.safelCalllback(fallback, BREAKER_OPEN_ERROR)
			return nil
		}
		runErr := run()
		if runErr != nil {
			this.fail()
			this.safelCalllback(fallback, runErr)
			return runErr
		}
		this.success()
		return nil
	default:
		if err != nil {
			this.fail()
			this.safelCalllback(fallback, err)
			return err
		}
		this.success()
		return nil
	}
}

func Do(ctx context.Context, name string, run runFunc, fallback fallbackFunc) error {
	if run == nil {
		return FUNC_NIL_ERROR
	}
	if name == "" {
		return NAME_NIL_ERROR
	}
	pBreaker, err := getBreakerManager(name)
	if err != nil {
		fallback(err)
		return err
	}
	beforeDoErr := pBreaker.beforeDo(ctx, name)
	if beforeDoErr != nil {
		callBackErr := pBreaker.afterDo(ctx, run, fallback, beforeDoErr)
		return callBackErr
	}
	runErr := run()
	return pBreaker.afterDo(ctx, run, fallback, runErr)
}
