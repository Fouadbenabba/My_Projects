package controllers

import (
	"first/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCity(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Convert string to int
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid city ID")
	}
	city, err := database.GetCity(id)
	fmt.Println("city: ", city)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "City not found"})
	}
	return c.JSON(http.StatusOK, city)
}

func GetAllCities(c echo.Context) error {
	// Retrieve all cities from the database
	cities, err := database.GetAllCities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch cities"})
	}

	// Return the cities as JSON
	return c.JSON(http.StatusOK, cities)
}

type CreateCityRequest struct {
	Name    string `json:"name" validate:"required"`
	Country int    `json:"country" validate:"required"`
}

func CreateCity(c echo.Context) error {
	var req CreateCityRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println("req:::", req)
	err := database.CreateCity(req.Name, req.Country)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "City created successfully"})

}

func DeleteCity(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid city id ")
	}

	err = database.DeleteCity(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "city delet seccessfully"})

}

func UpdateCity(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid city id")
	}
	err = database.UpdateCity(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "city update seccessfully"})

}
