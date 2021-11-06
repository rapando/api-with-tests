// read this from
// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
// I used mysql instead

package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	a := App{}
	a.Initialize(os.Getenv("DB_URI"))
	a.Run()
}
