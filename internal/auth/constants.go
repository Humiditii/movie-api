package auth

type UserType int

const (
	User UserType = iota
	Admin
	Root
)
