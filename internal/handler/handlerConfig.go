package handler

import (
	"github.com/JulianKerns/butchers-order/internal/database"
)

type HandlerConfig struct {
	DB *database.Queries
}

func GetHandlerConfig() *HandlerConfig {

	return &HandlerConfig{}
}
