package user

import (
	"github.com/Unknwon/macaron"
)

const (
	TplSignIn string = "user/signin"
	TplSignUp string = "user/signup"
)

func SignUp(ctx *macaron.Context) {
	ctx.Data["Title"] = "注册"
	ctx.HTML(200, TplSignUp)
}

func SignIn(ctx *macaron.Context) {
	ctx.Data["Title"] = "登录"
	ctx.HTML(200, TplSignIn)
}

func SignUpPost(ctx *macaron.Context, form SignUpForm) {
	ctx.Data["Username"] = form.Username
	ctx.Data["Password"] = form.Password
	ctx.Data["Flash"] = "Signup Succeed!"

}

func SignInPost(ctx *macaron.Context, form SignInForm) {
	ctx.Data["Email"] = form.Email
	ctx.Data["Password"] = form.Password
	ctx.Data["Flash"] = "Signin Succeed!"

}
