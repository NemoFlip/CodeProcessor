package entity

type User struct {
	ID       string `json:"-"`
	Login    string `json:"username"`
	Password string `json:"password"`
}
