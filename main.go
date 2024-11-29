package main

import (
	_ "gf-demo-takeaway/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-demo-takeaway/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
