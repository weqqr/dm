package userproxy

import (
	"dm/internal/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type nameRequestBody struct {
	Name string `json:"name"`
}

func (g *Gateway) create(c *gin.Context) {
	var requestBody nameRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if len(requestBody.Name) == 0 {
		log.Print("Error: Field name is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Error: Field name is empty",
		})
		return
	}

	response, err := g.client.CreateUser(g.ctx, &domain.CreateUserRequest{
		Name: requestBody.Name,
	})

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      response.Id,
		"name":    requestBody.Name,
		"message": "Successful created user.",
	})
}

func (g *Gateway) get(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		log.Print("Error: Field id is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})
		return
	}

	response, err := g.client.GetUser(g.ctx, &domain.GetUserRequest{
		Id: id,
	})

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      response.User.Id,
		"name":    response.User.Name,
		"message": "Successful get data about user from database.",
	})
}

func (g *Gateway) update(c *gin.Context) {
	var requestBody nameRequestBody
	id := c.Param("id")

	if err := c.BindJSON(&requestBody); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if len(id) == 0 {
		log.Print("Error: Field id is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})
		return
	}

	if len(requestBody.Name) == 0 {
		log.Print("name is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Warning: Can't change, because name is empty",
		})
		return
	}

	_, err := g.client.UpdateUser(g.ctx, &domain.UpdateUserRequest{
		Id:   id,
		Name: &requestBody.Name,
	})

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"name":    requestBody.Name,
		"message": "Successful update data user from database.",
	})
}

func (g *Gateway) delete(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		log.Print("Error: Field id is empty")
		c.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})
		return
	}

	_, err := g.client.DeleteUser(g.ctx, &domain.DeleteUserRequest{
		Id: id,
	})

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful delete user from database.",
	})
}
