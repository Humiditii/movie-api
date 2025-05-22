package auth

type SigninDto struct {
	Email string
	Password string
}

type SignupDto struct {
	SigninDto
	UserType UserType
}




