package models

// Product
type Product struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Store    string `db:"store"`
	Stock    int    `db:"stock"`
	Category string `db:"category"`
}

var (
	ProductSelectStatement = `select * from product`
	ProductWhere           = `where stock > 0`
)
