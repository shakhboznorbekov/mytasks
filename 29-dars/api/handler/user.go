package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/asadbekGo/golang_crud/models"
	"github.com/asadbekGo/golang_crud/storage"
)

func (h *HandlerV1) Create(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error while creating: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.Create(h.db, user)
	if err != nil {
		log.Printf("error while creating: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while creating"))
		return
	}

	userId, err := storage.GetById(h.db, id)
	if err != nil {
		log.Printf("error while getting by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while getting by id"))
		return
	}

	c.JSON(http.StatusOK, userId)
}

func (h *HandlerV1) Update(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error while updating: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := storage.Update(h.db, user)
	if err != nil {
		log.Printf("error while updating: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while updating"))
		return
	}

	userId, err := storage.GetById(h.db, user.Id)
	if err != nil {
		log.Printf("error while getting by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while getting by id"))
		return
	}

	fmt.Println(rowsAffected)

	c.JSON(http.StatusOK, userId)
}

func (h *HandlerV1) Delete(c *gin.Context) {
	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error while updating: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = storage.Delete(h.db, user.Id)
	if err != nil {
		log.Printf("error while deleteing: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while deleting"))
		return
	}
}

func (h *HandlerV1) Patch(c *gin.Context) {
	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error while patching: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := storage.Patch(h.db, user)
	if err != nil {
		log.Printf("error whiling patch: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while patching"))
		return
	}

	userId, err := storage.GetById(h.db, user.Id)
	if err != nil {
		log.Printf("error while getting by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while patching"))
		return
	}

	fmt.Println(rowsAffected)

	c.JSON(http.StatusOK, userId)
}

func (h *HandlerV1) GetById(c *gin.Context) {

	id := c.Param("id")

	user, err := storage.GetById(h.db, id)
	if err != nil {
		log.Printf("error while getting by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while getting by id"))
		return
	}

	c.JSON(http.StatusOK, user)
}
func (h *HandlerV1) GetList(c *gin.Context) {

	users, err := storage.GetList(h.db)
	if err != nil {
		log.Printf("error while getting list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error while getting list"))
		return
	}

	c.JSON(http.StatusOK, users)
}
