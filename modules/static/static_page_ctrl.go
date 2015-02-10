package static

import (
	"github.com/Unknwon/macaron"
)

const (
	TplHomePage    string = "static_pages/home"
	TplAboutPage   string = "static_pages/about"
	TplContactPage string = "static_pages/contact"
	Tpl404Page     string = "static_pages/404"
)

func HomePage(ctx *macaron.Context) {
	ctx.Data["AppName"] = ctx.Tr("app.name")
	ctx.Data["Desc"] = ctx.Tr("app.desc")
	ctx.Data["Title"] = ctx.Tr("page.home")

	ctx.Data["IsHomePage"] = true
	ctx.HTML(200, TplHomePage)
}

func AboutPage(ctx *macaron.Context) {
	ctx.Data["AppName"] = ctx.Tr("app.name")
	ctx.Data["Desc"] = ctx.Tr("app.desc")
	ctx.Data["Title"] = ctx.Tr("page.about")
	ctx.Data["IsAboutPage"] = true
	ctx.HTML(200, TplAboutPage)
}

func ContactPage(ctx *macaron.Context) {
	ctx.Data["AppName"] = ctx.Tr("app.name")
	ctx.Data["Desc"] = ctx.Tr("app.desc")
	ctx.Data["Title"] = ctx.Tr("page.contact")
	ctx.Data["IsContactPage"] = true
	ctx.HTML(200, TplContactPage)
}

func NotFound(ctx *macaron.Context) {
	ctx.Data["AppName"] = ctx.Tr("app.name")
	ctx.Data["Desc"] = ctx.Tr("app.desc")
	ctx.Data["Title"] = ctx.Tr("page.not_found")
	ctx.HTML(404, Tpl404Page)
}
