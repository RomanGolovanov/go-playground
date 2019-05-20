package storage

import (
	"database/sql"
	"encoding/json"
	"go-playground/model"

	// load SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

// SqliteDeskStorage stores  desks in SQLite database
type SqliteDeskStorage struct {
	db     *sql.DB
	insert *sql.Stmt
	update *sql.Stmt
	delete *sql.Stmt
	get    *sql.Stmt
	getAll *sql.Stmt
}

// NewDesk adds new desk to repo
func (storage SqliteDeskStorage) NewDesk(desk model.Desk) bool {

	b, err := json.Marshal(desk)
	if err != nil {
		return false
	}

	js := string(b)

	_, ok := storage.GetDesk(desk.Name)

	if !ok {
		_, err := storage.insert.Exec(desk.Name, js)
		panicIfError(err)
	} else {
		_, err := storage.update.Exec(js, desk.Name)
		panicIfError(err)
	}

	return true
}

// GetDesk returns desk by name
func (storage SqliteDeskStorage) GetDesk(name string) (model.Desk, bool) {
	rows, err := storage.get.Query(name)
	panicIfError(err)

	defer rows.Close()

	if !rows.Next() {
		return model.Desk{}, false
	}

	var js string
	rows.Scan(&js)

	var desk model.Desk
	panicIfError(json.Unmarshal([]byte(js), &desk))

	return desk, true
}

// DeleteDesk deletes desk by name
func (storage SqliteDeskStorage) DeleteDesk(name string) bool {
	_, err := storage.delete.Exec(name)
	return err != nil
}

// GetAllDesks return all desks
func (storage SqliteDeskStorage) GetAllDesks() []model.Desk {
	rows, err := storage.getAll.Query()
	panicIfError(err)

	defer rows.Close()

	desks := make([]model.Desk, 0)
	for rows.Next() {
		var js string
		rows.Scan(&js)
		var desk model.Desk
		panicIfError(json.Unmarshal([]byte(js), &desk))
		desks = append(desks, desk)
	}

	return desks

}

func (storage SqliteDeskStorage) Close() {
	storage.db.Close()
}

// CreateSqliteDeskStorage creates storage using fileName
func CreateSqliteDeskStorage(fileName string) SqliteDeskStorage {

	db, err := sql.Open("sqlite3", fileName)
	panicIfError(err)

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS desks (name TEXT PRIMARY KEY NOT NULL, json TEXT NOT NULL)")
	panicIfError(err)

	_, err = statement.Exec()
	panicIfError(err)

	insert, err := db.Prepare("INSERT INTO  desks (name, json) VALUES (?, ?)")
	panicIfError(err)

	update, err := db.Prepare("UPDATE desks SET json = ? WHERE name = ?")
	panicIfError(err)

	delete, err := db.Prepare("DELETE FROM desks WHERE name = ?")
	panicIfError(err)

	get, err := db.Prepare("SELECT json FROM desks WHERE name = ?")
	panicIfError(err)

	getAll, err := db.Prepare("SELECT json FROM desks")
	panicIfError(err)

	storage := SqliteDeskStorage{db, insert, update, delete, get, getAll}

	return storage
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
