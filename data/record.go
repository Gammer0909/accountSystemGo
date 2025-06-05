package data

type Record struct {
	Id           int    `csv:"id"`
	Username     string `csv:"username"`
	Email        string `csv:"email"`
	PasswordHash string `csv:"password"`
}
