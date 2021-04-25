package repository

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
	"practice/models"
)


func GetPerson(id int) (*models.Person, error) {
	db, err := sql.Open("mysql", "root:heyaoxu555@/test")
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("DB connect error "))
		return nil, err
	}
	defer db.Close()
	rows := db.QueryRow("SELECT id, name, age FROM Person WHERE id = ?", id)
	var person = models.Person{}
	err = rows.Scan(&person.Id, &person.Name, &person.Age)
	if err != nil {
		return nil, nil
	}
	return &person, nil
}


