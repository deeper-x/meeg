package db

import "database/sql"

type Instance struct {
	db *sql.DB
}

func newConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "migrations.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewDB instance builder
func NewDB(db *sql.DB) (Instance, error) {
	return Instance{
		db: db,
	}, nil
}

// Select is a fetching wrapper
func (i Instance) Select(query string, id int) ([]string, error) {
	var res []string
	rows, err := i.db.Query(query, id)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return res, err
		}

		res = append(res, name)
	}
	return res, nil
}
