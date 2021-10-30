package cmd

import (
	"LinuxTools/router"
	"github.com/urfave/cli/v2"
)

func zeroFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "查看系统信息",
		},
		&cli.BoolFlag{
			Name:    "replace",
			Aliases: []string{"rep"},
			Usage:   "更换系统源",
		},
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "查询端口占用",
		},
	}
}

func zeroAction(c *cli.Context) error {
	for _, item := range zeroFlags() {
		if len(c.String(item.Names()[0])) != 0 || len(c.String(item.Names()[1])) != 0 {
			router.RouterInit(item.Names(), c)
		}
	}
	return nil
}

func Cmd() *cli.App {
	app := &cli.App{
		Name:      "LinuxTools",
		Usage:     "Tools",
		UsageText: "Linux常用工具箱",
		Action:    zeroAction,
		Flags:     zeroFlags(),
	}
	return app
}
