package main

import (
	"context"
	"os"

	"github.com/LyricTian/gin-admin/v6/internal/app"
	"github.com/LyricTian/gin-admin/v6/pkg/logger"
	"github.com/urfave/cli/v2"
)

var VERSION = "1.0.0"

func main() {
	logger.SetVersion(VERSION)
	ctx := logger.NewTraceIDContext(context.Background(), "main")

	appInstance := cli.NewApp()
	appInstance.Name = "Auth"
	appInstance.Version = VERSION
	appInstance.Usage = "RBAC scaffolding based on GIN + GORM + CASBIN + WIRE."
	appInstance.Commands = []*cli.Command{
		newWebCmd(ctx),
	}
	err := appInstance.Run(os.Args)
	if err != nil {
		logger.Errorf(ctx, err.Error())
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "运行web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.json,.yaml,.toml)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "model",
				Aliases:  []string{"m"},
				Usage:    "casbin的访问控制模型(.conf)",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "menu",
				Usage: "初始化菜单数据配置文件(.yaml)",
			},
			//&cli.StringFlag{
			//	Name:  "www",
			//	Usage: "静态站点目录",
			//},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetModelFile(c.String("model")),
				//app.SetWWWDir(c.String("www")),
				app.SetMenuFile(c.String("menu")),
				app.SetVersion(VERSION))
		},
	}
}
