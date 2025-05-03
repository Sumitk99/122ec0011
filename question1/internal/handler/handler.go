package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/Sumitk99/122ec0011/question1/internal/models"
	"github.com/Sumitk99/122ec0011/question1/internal/service"
	"github.com/gin-gonic/gin"
)


func Handler(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		numID := c.Param("numberid") 
		
		if numID == "p" {
			PrimeHandler(c, db)
		} else if numID == "f" {
			FibHandler(c, db)
		} else if numID == "e" {
			EvenHandler(c, db)
		} else if numID == "r" {
			RandomHandler(c, db)
		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error":"Invalid numberid"})
		}
	}	
}

func PrimeHandler(c *gin.Context, db *sql.DB){
	req, _ := http.NewRequest("GET", "http://20.244.56.144/evaluation-service/primes", nil)
	authToken := os.Getenv("token")
	req.Header.Set("authorization", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	input := &models.InputData{}
	if err := json.NewDecoder(resp.Body).Decode(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"can't parse input json from test server"})
		return
	}
	
	service.StoreUniqueNumbers(db)
	UpdatedNumbers := service.FetchLatestData(db)

	PrevStateSize := len(UpdatedNumbers) - len(input.Numbers)
	if PrevStateSize < 0 {
		PrevStateSize = 0
	}
	PrevState := UpdatedNumbers[:PrevStateSize]

	StartIndex := 0
	if len(UpdatedNumbers) > 10 {
		StartIndex = len(UpdatedNumbers) - 10
	}
	NewWindow := UpdatedNumbers[StartIndex:]

	NewAvg := service.CalculateAvg(UpdatedNumbers)
	c.JSON(http.StatusOK, models.Response{
		Numbers: UpdatedNumbers,
		WindowPrevState: PrevState,
		WindowCurrState: NewWindow,
		Avg: NewAvg,
	})
}

func FibHandler(c *gin.Context, db *sql.DB){
	req, _ := http.NewRequest("GET", "http://20.244.56.144/evaluation-service/fibo", nil)
	authToken := os.Getenv("token")
	req.Header.Set("authorization", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	input := &models.InputData{}
	if err := json.NewDecoder(resp.Body).Decode(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"can't parse input json from test server"})
		return
	}
	
	service.StoreUniqueNumbers(db)
	UpdatedNumbers := service.FetchLatestData(db)

	PrevStateSize := len(UpdatedNumbers) - len(input.Numbers)
	if PrevStateSize < 0 {
		PrevStateSize = 0
	}
	PrevState := UpdatedNumbers[:PrevStateSize]

	StartIndex := 0
	if len(UpdatedNumbers) > 10 {
		StartIndex = len(UpdatedNumbers) - 10
	}
	NewWindow := UpdatedNumbers[StartIndex:]

	NewAvg := service.CalculateAvg(UpdatedNumbers)
	c.JSON(http.StatusOK, models.Response{
		Numbers: UpdatedNumbers,
		WindowPrevState: PrevState,
		WindowCurrState: NewWindow,
		Avg: NewAvg,
	})

}

func EvenHandler(c *gin.Context, db *sql.DB){
	req, _ := http.NewRequest("GET", "http://20.244.56.144/evaluation-service/even", nil)
	authToken := os.Getenv("token")
	req.Header.Set("authorization", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	input := &models.InputData{}
	if err := json.NewDecoder(resp.Body).Decode(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"can't parse input json from test server"})
		return
	}
	
	service.StoreUniqueNumbers(db)
	UpdatedNumbers := service.FetchLatestData(db)

	PrevStateSize := len(UpdatedNumbers) - len(input.Numbers)
	if PrevStateSize < 0 {
		PrevStateSize = 0
	}
	PrevState := UpdatedNumbers[:PrevStateSize]

	StartIndex := 0
	if len(UpdatedNumbers) > 10 {
		StartIndex = len(UpdatedNumbers) - 10
	}
	NewWindow := UpdatedNumbers[StartIndex:]

	NewAvg := service.CalculateAvg(UpdatedNumbers)
	c.JSON(http.StatusOK, models.Response{
		Numbers: UpdatedNumbers,
		WindowPrevState: PrevState,
		WindowCurrState: NewWindow,
		Avg: NewAvg,
	})

}

func RandomHandler(c *gin.Context, db *sql.DB){
	req, _ := http.NewRequest("GET", "http://20.244.56.144/evaluation-service/rand", nil)
	authToken := os.Getenv("token")
	req.Header.Set("authorization", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	input := &models.InputData{}
	if err := json.NewDecoder(resp.Body).Decode(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"can't parse input json from test server"})
		return
	}
	
	service.StoreUniqueNumbers(db)
	UpdatedNumbers := service.FetchLatestData(db)

	PrevStateSize := len(UpdatedNumbers) - len(input.Numbers)
	if PrevStateSize < 0 {
		PrevStateSize = 0
	}
	PrevState := UpdatedNumbers[:PrevStateSize]

	StartIndex := 0
	if len(UpdatedNumbers) > 10 {
		StartIndex = len(UpdatedNumbers) - 10
	}
	NewWindow := UpdatedNumbers[StartIndex:]

	NewAvg := service.CalculateAvg(UpdatedNumbers)
	c.JSON(http.StatusOK, models.Response{
		Numbers: UpdatedNumbers,
		WindowPrevState: PrevState,
		WindowCurrState: NewWindow,
		Avg: NewAvg,
	})

}