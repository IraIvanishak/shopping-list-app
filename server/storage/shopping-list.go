package storage

import (
	"database/sql"
	"log"

	"github.com/IraIvanishak/shopping-list-app/config"
)

type Good struct {
	ID     int    `josn:"id"`
	Name   string `json:"name"`
	Amount int    `josn:"amount"`
}

// GetAllGoods read all goods from the table
func GetAllGoods() ([]Good, error) {
	query := `SELECT name, amount FROM shopping_list`

	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Query execution error:", err)
		return nil, err
	}
	defer rows.Close()

	var goods []Good

	for rows.Next() {
		var good Good
		err := rows.Scan(&good.Name, &good.Amount)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		goods = append(goods, good)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	return goods, nil
}

// CreateGood inserts a new Good into the  table.
func CreateGood(good *Good) error {
	query := `INSERT INTO shopping_list (name, amount) VALUES ($1, $2) RETURNING id`

	err := config.DB.QueryRow(query, good.Name, good.Amount).Scan(&good.ID)
	if err != nil {
		log.Println("Error creating new Good:", err)
		return err
	}

	return nil
}

// UpdateGood updates an existing Good in the table.
func UpdateGood(good *Good) error {
	query := `UPDATE shopping_list SET name = $1, amount = $2 WHERE id = $3`

	res, err := config.DB.Exec(query, good.Name, good.Amount, good.ID)
	if err != nil {
		log.Println("Error updating Good:", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// DeleteGood removes a Good from the table.
func DeleteGood(id int) error {
	query := `DELETE FROM shopping_list WHERE id = $1`

	res, err := config.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting Good:", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
