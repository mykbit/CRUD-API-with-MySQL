package models

import (
	"database/sql"
	"errors"
)

type Vinyl struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	ReleaseDate string `json:"ReleaseDate"`
}

type VinylModel struct {
	DB *sql.DB
}

func (m *VinylModel) GetAll() ([]*Vinyl, error) {
	stmt := `SELECT id, title, description, releaseDate FROM vinyls 
	ORDER BY title ASC`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	vinyls := []*Vinyl{}

	for rows.Next() {
		v := &Vinyl{}

		err := rows.Scan(&v.ID, &v.Title, &v.Description, &v.ReleaseDate)
		if err != nil {
			return nil, err
		}

		vinyls = append(vinyls, v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return vinyls, nil
}

func (m *VinylModel) GetById(id int) (*Vinyl, error) {
	stmt := `SELECT id, title, description, releaseDate FROM vinyls
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	v := &Vinyl{}

	err := row.Scan(&v.ID, &v.Title, &v.Description, &v.ReleaseDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return v, nil
}

func (m *VinylModel) Insert(title string, description string, releaseDate string) (int, error) {
	stmt := `INSERT INTO vinyls (title, description, releaseDate) 
	VALUES (?, ?, ?)`

	result, err := m.DB.Exec(stmt, title, description, releaseDate)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *VinylModel) Update(id int, title string, description string, releaseDate string) error {
	stmt := `UPDATE vinyls SET title = ?, description = ?, releaseDate = ?
	WHERE id = ?`

	_, err := m.DB.Exec(stmt, title, description, releaseDate, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *VinylModel) Delete(id int) error {
	stmt := `DELETE FROM vinyls WHERE id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
