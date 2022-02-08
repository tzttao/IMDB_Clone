package main

import (
	Database "backend/Database"
	"backend/Router"
)

func main() {
	Router.Main()
	defer Database.DB.Close()
}
