package models

import (
	"database/sql"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllCategories(db *sql.DB, nameFilter string) ([]Category, error) {
	var rows *sql.Rows
	var err error

	if nameFilter != "" {
		rows, err = db.Query(`SELECT * FROM categories WHERE name LIKE ?`, "%"+nameFilter+"%")
	} else {
		rows, err = db.Query(`SELECT * FROM categories`)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		err = rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func CreateCategory(db *sql.DB, name string) error {
	_, err := db.Exec(`INSERT INTO categories(name) VALUES (?)`, name)
	return err
}

func UpdateCategory(db *sql.DB, id int, name string) error {
	_, err := db.Exec(`UPDATE categories SET name = ? WHERE id = ?`, name, id)
	return err
}

func DeleteCategory(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
