package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qu-bit1/spoTask/tree/main/task3/pkg/store"
	"log"
	"strconv"
)

func main() {

	// creating a database connection
	// Conn function made in store.go
	db := store.Conn()

	r := gin.Default()

	// making a post request
	r.POST("/signup", func(c *gin.Context) {

		id := c.PostForm("id")
		// post form returns a string value
		// so we convert string to int using parseint
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id = %v", i)
		password := c.PostForm("pwd")
		fmt.Printf("id: %d; password: %s", i, password)
		_, err = db.Exec("INSERT INTO users(userID,pwd) values(?,?)", i, password)
		if err != nil {
			log.Fatal(err)
		}

	})

	err := r.Run()
	if err != nil {
		return
	}

}
