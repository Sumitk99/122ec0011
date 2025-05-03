package models

type InputData struct {
	Numbers []int `json:"numbers"`
}

type DataFromDB struct {
	Numbers []int 
}

type Response struct {
	WindowPrevState []int `json:"windowPrevState"`
	WindowCurrState []int `json:"windowCurrState"`
	Numbers []int `json:"numbers"`
	Avg float64 `json:"avg"`
}