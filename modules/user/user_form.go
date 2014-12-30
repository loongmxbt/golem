package user

type SignUpForm struct {
	Username string `form:"uname" binding:"Required;AlphaDashDot;MaxSize(35)"`
	Email    string `form:"email" binding:"Required;Email;MaxSize(50)"`
	Password string `form:"passwd" binding:"Required;MinSize(6);MaxSize(255)"`
	Retype   string `form:"retype"`
}

type SignInForm struct {
	Email    string `form:"email" binding:"Required;Email;MaxSize(50)"`
	Password string `form:"passwd" binding:"Required;MinSize(6);MaxSize(255)"`
	Remember bool   `form:"remember"`
}
