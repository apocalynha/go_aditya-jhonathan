package controllers

import (
	"eksplorasi/config"
	"eksplorasi/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// CreateBlogController: Membuat blog baru dalam database
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// GetBlogsController: Mengambil semua blog dari database beserta informasi User terkait
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := config.DB.Preload("User").Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// GetBlogController: Mengambil blog berdasarkan ID beserta informasi User terkait
func GetBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog models.Blog
	if err := config.DB.Preload("User").First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by ID",
		"blog":    blog,
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
