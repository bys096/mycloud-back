package request

type JoinDto struct {
	Id         uint64
	Name       string
	Email      string
	Password   string
	PasswordCf string
}
