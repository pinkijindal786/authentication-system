package repositories

import (
	"Authentication_System/internal/models"
	"log"
	"os/exec"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var tokens = &models.JWTToken{
	Token: "dummy",
}

var users = &models.User{
	ID:       1,
	Password: "test",
	Email:    "test@gmail.com",
	IsActive: true,
}

func cleanup(dbName string) *gorm.DB {
	var err error
	// remove old database
	exec.Command("rm", "-f", dbName)

	// open and create a new database
	gdb, err := gorm.Open(sqlite.Open("./testDB/"+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// migrate tables
	gdb.AutoMigrate(&models.JWTToken{})

	// add mock data
	gdb.Create(tokens)

	// migrate tables
	gdb.AutoMigrate(&models.User{})

	// add mock data
	gdb.Create(users)
	return gdb.Debug().Begin()
}
