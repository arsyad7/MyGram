package main

import (
	"log"
	"mygram/database"

	"github.com/gin-gonic/gin"
)

const PORT = ":3636"

func main() {
	_ = database.StartDB()
	r := gin.Default()

	log.Println("Server is listening at port", PORT)
	r.Run(PORT)
}
