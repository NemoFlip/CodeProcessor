package entity

type User struct {
	ID       string `json:"id"`
	Login    string `json:"username"`
	Password string `json:"password"`
}
