package entity

type Task struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}
