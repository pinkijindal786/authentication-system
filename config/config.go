package config

import (
	"os"
)

var (
	JWTSecret = os.Getenv("JWT_SECRET")
	DBConnStr = os.Getenv("DB_CONN_STR")
)
