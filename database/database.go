package database

import (
	"database/sql"
	"first/api/models"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	fmt.Println("---", dbPassword)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db
}
func CreateCity(name string, country int) error {
	fmt.Println("db val :", db)
	_, err := db.Exec("INSERT INTO cities (name, country) VALUES ($1, $2)", name, country)
	if err != nil {
		return fmt.Errorf("could not insert city: %v", err)
	}
	return nil
}
func GetCity(id int) (models.City, error) {
	// Query to select the city by ID
	query := "SELECT id, name, country FROM cities WHERE id = $1"
	row := db.QueryRow(query, id)
	fmt.Printf("Row queried for ID %d\n", id)
	// Initialize a City struct to hold the result
	var city models.City
	err := row.Scan(&city.ID, &city.CityName, &city.CountryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return city, sql.ErrNoRows
		}
		return city, fmt.Errorf("failed to fetch city: %v", err)
	}
	return city, nil
}

func GetAllCities() ([]models.City, error) {
	query := "SELECT id, name, country FROM cities"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cities: %v", err)
	}
	defer rows.Close()

	var cities []models.City

	for rows.Next() {
		var city models.City
		err := rows.Scan(&city.ID, &city.CityName, &city.CountryID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan city: %v", err)
		}
		cities = append(cities, city)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return cities, nil
}

func DeleteCity(id int) error {
	query := "DELETE FROM cities WHERE id= $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete city: %v", err)
	}
	return nil
}

func UpdateCity(id int) error {
	query := "UPDATE cities SET country = '9' WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("city not updated:%v", err)

	}
	return nil
}
