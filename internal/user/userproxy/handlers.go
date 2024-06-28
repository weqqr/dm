package userproxy

import (
	"dm/internal/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func (g *Gateway) create(c *gin.Context) {
	name := c.Param("name")

	response, err := g.client.CreateUser(g.ctx, &domain.CreateUserRequest{
		Name: name,
	})

	if err != nil {
		log.Printf("GATEWAY | USER | CREATE | Error: %v", err)
		return
	}

	if len(name) == 0 {
		log.Print("GATEWAY | USER | CREATE | Warning: Can't create, because name is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Warning: Can't create, because name is empty",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      response.Id,
		"name":    name,
		"message": "Successful created user.",
	})
}

func (g *Gateway) get(c *gin.Context) {
	id := c.Param("id")

	response, err := g.client.GetUser(g.ctx, &domain.GetUserRequest{
		Id: id,
	})

	if err != nil {
		log.Printf("GATEWAY | USER | GET | Error: %v", err)
		return
	}

	if len(id) == 0 {
		log.Print("GATEWAY | USER | GET | Error: Field id is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      response.User.Id,
		"name":    response.User.Name,
		"message": "Successful get data about user from database.",
	})
}

func (g *Gateway) change(c *gin.Context) {
	id, name := c.Param("id"), c.Param("name")

	_, err := g.client.UpdateUser(g.ctx, &domain.UpdateUserRequest{
		Id:   id,
		Name: &name,
	})

	if err != nil {
		log.Printf("GATEWAY | USER | GET | Error: %v", err)
		return
	}

	if len(id) == 0 {
		log.Print("GATEWAY | DELETE | Error: Field id is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})
		return
	}

	if len(name) == 0 {
		log.Print("GATEWAY | USER | GET | Warning: Can't change, because name is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Warning: Can't change, because name is empty",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"name":    name,
		"message": "Successful update data user from database.",
	})
}

func (g *Gateway) delete(c *gin.Context) {
	id := c.Param("id")
	_, err := g.client.DeleteUser(g.ctx, &domain.DeleteUserRequest{
		Id: id,
	})

	if err != nil {
		log.Printf("GATEWAY | DELETE | Error: %v", err)
		return
	}

	if len(id) == 0 {
		log.Print("GATEWAY | DELETE | Error: Field id is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful delete user from database.",
	})
}
