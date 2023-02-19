package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	query := "SELECT name FROM demo WHERE rowid = ?"
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"demo"}))

	inst, err := NewDB(db)
	if err != nil {
		t.Error(err)
	}

	_, err = inst.Select(query, 1)
	if err != nil {
		t.Error(err)
	}

}
