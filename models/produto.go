package models

import (
	"strconv"

	"github.com/kentoimayoshi/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	conn, err := db.ConectaComBancoDeDados()
	if err != nil {
		return []Produto{}
	}
	defer conn.Close()

	rows, err := conn.Query("select id, nome, descricao, preco, quantidade from produtos order by id")
	if err != nil {
		return []Produto{}
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var p Produto
		if err = rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade); err != nil {
			return []Produto{}
		}
		produtos = append(produtos, p)
	}
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	conn, err := db.ConectaComBancoDeDados()
	if err != nil {
		return
	}
	defer conn.Close()

	_, _ = conn.Exec(`insert into produtos (nome, descricao, preco, quantidade) values ($1,$2,$3,$4)`,
		nome, descricao, preco, quantidade)
}

func DeletaProduto(id string) {
	conn, err := db.ConectaComBancoDeDados()
	if err != nil {
		return
	}
	defer conn.Close()

	_, _ = conn.Exec(`delete from produtos where id = $1`, id)
}

func EditaProduto(id string) Produto {
	conn, err := db.ConectaComBancoDeDados()
	if err != nil {
		return Produto{}
	}
	defer conn.Close()

	var p Produto
	_ = conn.QueryRow(`select id, nome, descricao, preco, quantidade from produtos where id=$1`, id).
		Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
	return p
}

func AtualizaProduto(id, nome, descricao, precoStr, qtdStr string) {
	conn, err := db.ConectaComBancoDeDados()
	if err != nil {
		return
	}
	defer conn.Close()

	preco, _ := strconv.ParseFloat(precoStr, 64)
	qtd, _ := strconv.Atoi(qtdStr)
	_, _ = conn.Exec(`update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5`,
		nome, descricao, preco, qtd, id)
}
