package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-shafaq/defcase"
	gg "github.com/go-shafaq/gin"
	"github.com/ruziba3vich/databaseLesson/internal/handlers"
)

func main() {
	router := gin.Default()

	dcase := defcase.Get()
	dcase.SetCase("json", "*", defcase.Snak_case)
	dcase.SetCase("form", "*", defcase.Snak_case)

	gg.SetDefCase(dcase)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "userRequests")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		PrintError(err)
	}
	defer db.Close()

	dbNames := []string{"users", "admins"}

	for _, dbName := range dbNames {
		name := "../internal/db/" + dbName + ".sql"
		sqlFile, err := os.ReadFile(name)
		if err != nil {
			PrintError(err)
		}

		_, err = db.Exec(string(sqlFile))
		if err != nil {
			PrintError(err)
		}
	}

	router.POST("/create-person", func(c *gin.Context) {
		handlers.CreatePersonHandler(c, db)
	})

	router.DELETE("/delete-person/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		handlers.DeletePersonHandler(c, id, db)
	})

	router.GET("/get-all-persons", func (c *gin.Context) {
		handlers.GetAllPersons(c, db)
	})

	router.GET("/average-age", func (c *gin.Context) {
		handlers.GetAverageAgeHandler(c, db)
	})

	router.GET("/count-persons-department", func (c *gin.Context) {
		handlers.GetCountOfPersonsInEachDepartment(c, db)
	})

	router.GET("/eldest-five-persons", func (c *gin.Context) {
		handlers.GetEldestFivePersons(c, db)
	})

	router.GET("/name-and-department", func (c *gin.Context) {
		handlers.GetNameAndDepartmentOnly(c, db)
	})

	router.GET("/older-than-thirty", func (c *gin.Context) {
		handlers.GetOlderThanThirty(c, db)
	})

	router.GET("/people-increasing-order", func (c *gin.Context) {
		handlers.GetPeopleInIncreasingOrder(c, db)
	})

	router.GET("/persons-starting-a", func (c *gin.Context) {
		handlers.GetPeopleStartingWithA(c, db)
	})

	router.GET("/twenty-five-thirty-five", func (c *gin.Context) {
		handlers.GetPeopleFromTwentyFiveToThirtyFive(c, db)
	})

	router.PUT("/update-department/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		handlers.UpdateDepartmentById(c, id, db)
	})

	router.PUT("/update-person-age/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		handlers.UpdatePersonByAge(c, id, db)
	})

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := router.Run(address); err != nil {
		PrintError(err)
	}
}

func PrintError(err error) {
	log.Fatal("Error :", err)
}
