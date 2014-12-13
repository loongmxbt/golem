package static

import (
	"github.com/Unknwon/macaron"
)

func HomePage(ctx *macaron.Context) {
	ctx.HTML(200, "static_pages/home")
}

func AboutPage(ctx *macaron.Context) {
	ctx.HTML(200, "static_pages/about")
}
