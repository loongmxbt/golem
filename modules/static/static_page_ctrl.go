package static

import (
	"github.com/Unknwon/macaron"
)

const (
	TplHomePage  string = "static_pages/home"
	TplAboutPage string = "static_pages/about"
)

func HomePage(ctx *macaron.Context) {
	ctx.Data["Title"] = "主页"
	ctx.Data["IsHomePage"] = true
	ctx.HTML(200, TplHomePage)
}

func AboutPage(ctx *macaron.Context) {
	ctx.Data["Title"] = "关于"
	ctx.Data["IsAboutPage"] = true
	ctx.HTML(200, TplAboutPage)
}
