package router

import (
	"LinuxTools/controller"
	"github.com/urfave/cli/v2"
)

func RouterInit(routerList []string, c *cli.Context) {
	Router(routerList, c)
}

func Router(routerList []string, c *cli.Context) {
	data := c.String(routerList[0])
	switch routerList[0] {
	case "version":
		controller.Version(data)
	case "replace":
		controller.ReplaceOsSource(data)
	case "port":
		controller.FindPort(data)
	}
}
