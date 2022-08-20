package controllers

import (
	
	"time"
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/Rei-Suzuki1729/spajam_stech_api/models"
	"github.com/Rei-Suzuki1729/spajam_stech_api/database"
)

//----------
// Handlers
//----------
var db = database.Connect()

func CreatePost(c echo.Context) error {
	
	newPost := models.Post{}
	
	err := c.Bind(&newPost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&newPost)
	return c.JSON(http.StatusCreated, newPost)
	
}

func GetPostById(c echo.Context) error {

	post := models.Post{}
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.First(&post, id)
	return c.JSON(http.StatusOK, post)
}

func GetPosts(c echo.Context) error {

	postList:= []models.Post{}
	db.Find(&postList)
	return c.JSON(http.StatusOK, postList)

}

func UpdatePost(c echo.Context) error {

	newPost := models.Post{}

	err := c.Bind(&newPost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	post := models.Post{}
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&post, id)

	post.Name = newPost.Name
	post.Time = newPost.Time
	post.ImageUrl = newPost.ImageUrl
	post.UpdatedAt = time.Now()

	db.Save(&post)
	return c.JSON(http.StatusOK, post)

}

func DeletePost(c echo.Context) error {

	post := models.Post{}
	id, _ := strconv.Atoi(c.Param("id"))
	
	err := c.Bind(&post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Delete(&post, id)

	return c.JSON(http.StatusNoContent, post)
}
