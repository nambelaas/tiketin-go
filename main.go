package main

import (
	"github.com/tiketin-management-api-with-go/config"
	"github.com/tiketin-management-api-with-go/database"
)

func main() {
	config.Init()
	database.Init()
}
