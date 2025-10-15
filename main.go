package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	_ = godotenv.Load()
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func conectaComBancoDeDados() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, fmtError("DATABASE_URL não definida")
	}
	return sql.Open("postgres", url)
}

type envErr string

func (e envErr) Error() string { return string(e) }

func fmtError(msg string) error { return envErr(msg) }

func index(w http.ResponseWriter, r *http.Request) {
	db, err := conectaComBancoDeDados()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var produtos []Produto

	for rows.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		if err = rows.Scan(&id, &nome, &descricao, &preco, &quantidade); err != nil {
			panic(err.Error())
		}

		p := Produto{
			Id:         id,
			Nome:       nome,
			Descricao:  descricao,
			Preco:      preco,
			Quantidade: quantidade,
		}
		produtos = append(produtos, p)
	}

	_ = tmpl.ExecuteTemplate(w, "Index", produtos)
}
