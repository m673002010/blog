package param

type AddUserParam struct {
	Username string `form:"username" binding:"required,min=3,max=10"`
	Password string `form:"password" binding:"required,min=6,max=24"`
	Account  string `form:"account" binding:"required,min=3,max=18"`
}

type EditUserParam struct {
	Id       int    `form:"id" binding:"required"`
	Username string `form:"username" binding:"min=3,max=10"`
	Account  string `form:"account" binding:"min=3,max=18"`
}

type LoginParam struct {
	Account  string `form:"account" binding:"required,min=3,max=18"`
	Password string `form:"password" binding:"required"`
}
