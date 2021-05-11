package repository

import (
	"errors"
	"fmt"
	"github.com/prakasitnan/go-template/src/config"
	"github.com/prakasitnan/go-template/src/model"
)

func GetAllUsers() ([]model.User, error){

	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int64) (model.User, error) {
	user := model.User{}
	row := config.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)

	err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Status)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user model.User) (model.User, error) {
	fmt.Println(user)
	res, err := config.DB.Exec("INSERT INTO users (name, username, password, status) VALUES ($1, $2, $3, $4)", "test2", "test2", "test2", 1)
	if err != nil {
		return user, errors.New("500. Internal Server Error." + err.Error())
	}

	id, errLastId := res.LastInsertId()
	if errLastId != nil {
		return user, errors.New("test "+ errLastId.Error())
	}

	user.Id = id
	return user, nil
}

func UpdateUser(id int64, user model.User) (model.User, error) {
	_, err := config.DB.Exec("UPDATE user SET name = $2, username=$3, password=$4, status=$5 WHERE id = $1", id, user.Name, user.Password, user.Status)
	if err != nil {
		return user, err
	}
	return user, err
}

func DeleteUser(id int64) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = $1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}