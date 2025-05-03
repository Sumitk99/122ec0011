package service

import (
	"database/sql"

	"github.com/Sumitk99/122ec0011/question1/internal/config"
	"github.com/Sumitk99/122ec0011/question1/internal/models"
	"github.com/Sumitk99/122ec0011/question1/internal/repository"
)

func StoreUniqueNumbers(numbers *models.InputData, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO numbers (value) VALUES ($1)", 42)
	if err != nil {
		return err
	}



}

func FetchLatestData(db *sql.DB) []int{
	repository.GetNumbersFromDB(db)
}

func CalculateAvg(numbers []int) float64 {
	sum := 0
	count := 0

	for i := len(numbers) - 1; i >= 0 && count < config.WINDOW_SIZE; i-- {
		sum += numbers[i]
		count++
	}

	avg := float64(sum) / float64(count)
	return avg
}