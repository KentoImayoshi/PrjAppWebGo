package routes

import (
	"net/http"

	"github.com/kentoimayoshi/controllers"
)

func Init() {
	http.HandleFunc("/", controllers.Index)
}
