package products

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vbrenister/go-rest-example/pkg/model"
)

type Repo interface {
	GetAll() ([]model.Product, error)
}

type productsRepo struct {
	db *sql.DB
}

func NewRepo() (*productsRepo, error) {
	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		return nil, err
	}
	return &productsRepo{
		db: db,
	}, nil
}


func (p *productsRepo) GetAll() ([]model.Product, error) {
	var produts []model.Product
	rows, err := p.db.Query("select id, name, inventory, price from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var pr model.Product

		if err := rows.Scan(&pr.ID, &pr.Name, &pr.Inventory, &pr.Price); err != nil {
			return nil, err
		}
		produts = append(produts, pr)
	}

	return produts, nil
}
