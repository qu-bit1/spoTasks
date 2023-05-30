package main

import (
	"database/sql"
	"log"
)

func main() {

	type User struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	func main() {
		router := gin.Default()

		// Set up MySQL connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/signup")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Create users table if it doesn't exist
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE,
		password VARCHAR(100)
	)`)
		if err != nil {
			log.Fatal(err)
		}

		// Handle signup request
		router.POST("/signup", func(c *gin.Context) {
			var user User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(400, gin.H{"error": "Invalid request"})
				return
			}

			// Encrypt the password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(500, gin.H{"error": "Internal server error"})
				return
			}

			// Insert user into the database
			_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, hashedPassword)
			if err != nil {
				if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
					// Error number 1062 indicates a duplicate entry error (username already exists)
					c.JSON(409, gin.H{"error": "Username already exists"})
					return
				}
				c.JSON(500, gin.H{"error": "Internal server error"})
				return
			}

			c.JSON(200, gin.H{"message": "Signup successful"})
		})

		// Start the server
		router.Run(":8080")
	}

}
