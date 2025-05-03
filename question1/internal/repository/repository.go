package repository

import (
	"database/sql"

	"github.com/Sumitk99/122ec0011/question1/internal/models"
	_ "github.com/lib/pq"
)
func ConnectToPostgres(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func StoreNumbersInPostgres(db *sql.DB, numbers *models.InputData) error {
	for _, number := range numbers.Numbers {
		_, err := db.Exec(`INSERT INTO numbers (value) VALUES ($1) ON CONFLICT (value) DO NOTHING`, number)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetNumbersFromDB(db *sql.DB)([]int, error){

	rows, err := db.Query("SELECT value FROM numbers ORDER BY created_at ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allNumbers []int
	for rows.Next() {
		var value int
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		allNumbers = append(allNumbers, value)
	}

	return allNumbers, nil

}