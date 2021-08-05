package usecase

import (
	"database/sql"
)

type ContractUC struct {
	ReqID     string
	UserID    string
	EnvConfig map[string]string
	DB        *sql.DB
	TX        *sql.DB
}
