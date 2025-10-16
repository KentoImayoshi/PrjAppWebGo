package controllers

import (
	"html/template"
	"net/http"

	"github.com/kentoimayoshi/db"
	"github.com/kentoimayoshi/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	rows, err := conn.Query("select id, nome, descricao, preco, quantidade from produtos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []models.Produto
	for rows.Next() {
		var p models.Produto
		if err = rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, p)
	}

	_ = tmpl.ExecuteTemplate(w, "Index", produtos)
}
