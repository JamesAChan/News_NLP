package main

import (
	"public-data/implement/google"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main()  {
	errL := godotenv.Load()
	if errL != nil {
		panic("Error loading .env file")
	}
	api := os.Getenv("API_GOOGLENEWS")
	res := google.Query(24, api)
	fmt.Println(res)
}
