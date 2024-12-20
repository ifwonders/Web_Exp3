package cmd

import (
	"context"
	"gf-demo-takeaway/internal/controller/takeaway"
	"gf-demo-takeaway/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of takeaway api",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				var (
					takeawayCtrl = takeaway.NewV2()
				)

				group.Bind(takeawayCtrl)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().AuthMiddleware)
					group.ALLMap(g.Map{
						"/protected/reset_password": takeawayCtrl.CustomerResetPassword,
					})

				})
			})
			s.Run()
			return nil
		},
	}
)
