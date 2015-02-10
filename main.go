package main

import (
	"github.com/Unknwon/macaron"
	"github.com/loongmxbt/golem/modules/base"
	"github.com/loongmxbt/golem/modules/config"
	"github.com/loongmxbt/golem/modules/static"
	"github.com/loongmxbt/golem/modules/user"
	"github.com/macaron-contrib/binding"
	"github.com/macaron-contrib/i18n"
)

func main() {

	base.GlobalInit()

	// Classic includes Logger, Recovery, Static
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(i18n.I18n(i18n.Options{
		Directory: "config/locale",
		Langs:     config.Langs,
		Names:     config.Names,
		Redirect:  true,
	}))

	// Routers
	// StaticPage routers
	m.Get("/", static.HomePage)
	m.Get("/about", static.AboutPage)
	m.Get("/contact", static.ContactPage)

	// User routers
	m.Group("/user", func() {
		m.Get("/signup", user.SignUp)
		m.Post("/signup", binding.Bind(user.SignUpForm{}), user.SignUpPost)
		m.Get("/signin", user.SignIn)
		m.Post("/signin", binding.Bind(user.SignInForm{}), user.SignInPost)
		m.Get("/profile", user.Profile)
	})

	// Admin routers

	// Node routers
	// m.Group("/node", func() {
	// 	m.Get("/:id", GetNode)
	// 	m.Post("/new", NewNode)
	// 	m.Put("/update/:id", UpdateNode)
	// 	m.Delete("/delete/:id", DeleteNode)
	// })

	m.Run()
}
