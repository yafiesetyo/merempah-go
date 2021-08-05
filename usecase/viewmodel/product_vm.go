package viewmodel

type ProductVM struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Store    string `json:"store"`
	Stock    int    `json:"stock"`
	Category string `json:"category"`
}
