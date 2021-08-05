package bootstrap

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/yafiesetyo/merempah-api-clone/usecase"
)

type Bootstrap struct {
	DB         *sql.DB
	App        *fiber.App
	ContractUC usecase.ContractUC
}
