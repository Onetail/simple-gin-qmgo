package controller

import (
	"fmt"
	"simple-gin-api/app"
	"strconv"
	"time"

	model "simple-gin-api/app/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	app    *app.App
	server *app.HTTPServer
	Router *gin.RouterGroup
}

const UserCollection = "user"

func (u *User) Init(server *app.HTTPServer) {

	u.server = server
	u.app = server.App
	u.Router = server.GetEngine().Group("/api/v1/user")

	u.Router.GET("/list", u.GetAllUser)
}

// Get All User Endpoint
func (u *User) GetAllUser(c *gin.Context) {
	db, err := u.app.Database.Init()
	if err != nil {
		fmt.Println(err)
	}

	users := model.Users{}
	err = db.Find(c, bson.M{}).All(&users)

	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"user": &users,
	})
}

// Get User Endpoint
func (u *User) GetUser(c *gin.Context) {
	db, err := u.app.Database.Init()
	if err != nil {
		fmt.Println(err)
	}

	id := c.Param("id")                   // Get Param
	idParse, errParse := strconv.Atoi(id) // Convert String to Int
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	user := model.User{}
	err = db.Find(c, bson.M{"id": &idParse}).One(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get User",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": &user,
	})
}

// Create User Endpoint
func (u *User) CreateUser(c *gin.Context) {
	db, err := u.app.Database.Init()
	if err != nil {
		fmt.Println(err)
	}

	user := model.User{}
	err = c.Bind(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = db.InsertOne(c, user)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Insert User",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succes Insert User",
		"user":    &user,
	})
}

// Update User Endpoint
func (u *User) UpdateUser(c *gin.Context) {
	db, err := u.app.Database.Init()
	if err != nil {
		fmt.Println(err)
	}

	id := c.Param("id")                   // Get Param
	idParse, errParse := strconv.Atoi(id) // Convert String to Int
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	user := model.User{}
	err = c.Bind(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	user.ID = idParse
	user.UpdatedAt = time.Now()

	err = db.UpdateOne(c, bson.M{"id": &idParse}, user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Update User",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succes Update User",
		"user":    &user,
	})
}

// Delete User Endpoint
func (u *User) DeleteUser(c *gin.Context) {
	db, err := u.app.Database.Init()
	if err != nil {
		fmt.Println(err)
	}

	id := c.Param("id")                   // Get Param
	idParse, errParse := strconv.Atoi(id) // Convert String to Int
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	err = db.Remove(c, bson.M{"id": &idParse})
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Delete User",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succes Delete User",
	})
}
