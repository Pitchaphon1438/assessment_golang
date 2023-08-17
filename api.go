package main

import (
	"example/go-orm-api/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/users", func(c *gin.Context) {
		var users []model.Users
		db.Find(&users)
		c.JSON(200, users)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.Users
		db.First(&user, id)
		data := map[string]interface{}{
			"ID":         int(user.ID),
			"Fname":      user.Fname,
			"Email":      user.Email,
			"Address":    user.Address,
			"Province":   user.Province,
			"PostalCode": user.PostalCode,
			"Country":    user.Country,
			"Phone":      user.Phone,
		}
		c.JSON(200, gin.H{"data": data})
	})
	r.POST("/user", func(c *gin.Context) {
		var user model.Users
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&user)
		c.JSON(200, gin.H{"RowsAffected": result.RowsAffected})
	})
	r.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.Users
		db.First(&user, id)
		db.Delete(&user)
		c.JSON(200, user)
	})
	r.PUT("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.Users
		var updateUser model.Users
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.First(&updateUser, id)
		updateUser.Username = user.Username
		updateUser.Fname = user.Fname
		updateUser.Email = user.Email
		updateUser.Address = user.Address
		updateUser.Province = user.Province
		updateUser.PostalCode = user.PostalCode
		updateUser.Country = user.Country
		updateUser.Phone = user.Phone
		db.Save(updateUser)
		c.JSON(200, updateUser)
	})
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
