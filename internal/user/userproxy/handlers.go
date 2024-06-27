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

func (proxy *UserProxy) create(ctx *gin.Context) {
	var requestBody nameRequestBody

	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Error: %v", err)

		return
	}

	if len(requestBody.Name) == 0 {
		log.Print("Error: Field name is empty")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Error: Field name is empty",
		})

		return
	}

	response, err := proxy.Client.CreateUser(ctx.Request.Context(), &domain.CreateUserRequest{
		Name: requestBody.Name,
	})

	if err != nil {
		log.Printf("Error: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":      response.GetId(),
		"name":    requestBody.Name,
		"message": "Successful created user.",
	})
}

func (proxy *UserProxy) get(ctx *gin.Context) {
	id := ctx.Param("id")

	if len(id) == 0 {
		log.Print("Error: Field id is empty")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})

		return
	}

	response, err := proxy.Client.GetUser(ctx.Request.Context(), &domain.GetUserRequest{
		Id: id,
	})

	if err != nil {
		log.Printf("Error: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":      response.GetUser().GetId(),
		"name":    response.GetUser().GetName(),
		"message": "Successful get data about user from database.",
	})
}

func (proxy *UserProxy) update(ctx *gin.Context) {
	var requestBody nameRequestBody
	id := ctx.Param("id")

	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Error: %v", err)

		return
	}

	if len(id) == 0 {
		log.Print("Error: Field id is empty")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})

		return
	}

	if len(requestBody.Name) == 0 {
		log.Print("name is empty")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Warning: Can't change, because name is empty",
		})

		return
	}

	_, err := proxy.Client.UpdateUser(ctx.Request.Context(), &domain.UpdateUserRequest{
		Id:   id,
		Name: &requestBody.Name,
	})

	if err != nil {
		log.Printf("Error: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":      id,
		"name":    requestBody.Name,
		"message": "Successful update data user from database.",
	})
}

func (proxy *UserProxy) delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if len(id) == 0 {
		log.Print("Error: Field id is empty")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Error: Field id is empty",
		})
		return
	}

	_, err := proxy.Client.DeleteUser(ctx.Request.Context(), &domain.DeleteUserRequest{
		Id: id,
	})

	if err != nil {
		log.Printf("Error: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successful delete user from database.",
	})
}
