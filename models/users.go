package models

import "apirest/db"

const Schema = `
	CREATE TABLE users(
		id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(30) NOT NULL,
		password VARCHAR(100) NOT NULL,
		email VARCHAR(30),
		create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
`

type User struct {
	Id       int64
	Name     string
	Password string
	Email    string
}

type Users []User

func NewUser(name, password, email string) *User {
	user := &User{
		Name:     name,
		Password: password,
		Email:    email,
	}
	return user
}

func CreateUser(name, password, email string) *User {
	user := NewUser(name, password, email)
	user.Save()
	return user
}

func (user *User) insert() {
	sql := "INSERT users SET name=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Name, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

func ListUsers() Users {
	sql := "SELECT id, name, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

func ListUserById(id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, name, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	}
	return user
}

func (user *User) Update() {
	sql := "UPDATE users SET name=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Name, user.Password, user.Email, user.Id)
}

func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.Update()
	}
}
