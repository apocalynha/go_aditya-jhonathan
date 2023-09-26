package controllers

import (
	"eksplorasi/config"
	"eksplorasi/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// Create Blog
func CreateBlogController(c echo.Context) error {
	// Parse the request body into a Blog struct
	var newBlog models.Blog
	if err := c.Bind(&newBlog); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	// Fetch the user based on the users_id
	var user models.User
	if err := config.DB.First(&user, newBlog.UsersID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid user ID",
			Data:    nil,
		})
	}

	// Associate the user with the blog
	newBlog.User = user

	// Save the blog post to the database
	result := config.DB.Create(&newBlog)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to create blog post",
			Data:    nil,
		})
	}

	// Return the created blog post as JSON response
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "Blog post created successfully",
		Data:    newBlog.ResponseConvert(),
	})
}

// Get All Blogs
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	result := config.DB.Preload("User").Find(&blogs)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to fetch blogs",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get all blogs",
		Data:    blogs,
	})
}

// Get blog by id
func GetBlogController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the blogs by their ID from the database using GORM
	var blog models.Blog
	result := config.DB.Preload("User").First(&blog, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, models.BaseResponse{
				Status:  false,
				Message: "ID Not Found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to fetch ID",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    blog,
	})
}

// UpdateBlogController: Memperbarui blog berdasarkan ID
func UpdateBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	c.Bind(&blog)
	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})
}

// DeleteBlogController: Menghapus blog berdasarkan ID
func DeleteBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	if err := config.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}
