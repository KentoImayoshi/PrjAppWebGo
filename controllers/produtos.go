package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/kentoimayoshi/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todos := models.BuscaTodosOsProdutos()
	_ = temp.ExecuteTemplate(w, "Index", todos)
}

func New(w http.ResponseWriter, r *http.Request) {
	_ = temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro preco:", err)
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeletaProduto(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.EditaProduto(id)
	_ = temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		models.AtualizaProduto(id, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
