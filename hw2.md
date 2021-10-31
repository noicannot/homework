作业：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


回答：
应该。
如果和其他库进行协作，考虑使用 errors.Wrap 或者 errors.Wrapf 保存堆栈信息。同样适用于和标准库协作的时候。 

返回代码：
return errors.Wrapf(err, fmt.Sprintf("sql err: %s ,err))

