package repository

import (
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