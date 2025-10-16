package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kentoimayoshi/routes"
)

func main() {
	_ = godotenv.Load()
	routes.CarregaRotas()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
