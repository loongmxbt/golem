package main

import (
	"github.com/Unknwon/macaron"
	"github.com/loongmxbt/golem/modules/static"
)

func main() {

	// Classic includes Logger, Recovery, Static
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	// Routers
	// StaticPage routers
	m.Get("/", static.HomePage)
	m.Get("/about", static.AboutPage)

	// User routers
	// m.Group("/user", func() {
	// 	m.Get("/signup", user.SignUp)
	// 	m.Post("/signup", binding.Bind(user.SignUpForm{}), user.SignUpPost)
	// 	m.Get("/signin", user.SignIn)
	// 	m.Post("/signin", binding.Bind(user.SignInForm{}), user.SignInPost)
	// 	m.Get("/profile", user.Profile)
	// })

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