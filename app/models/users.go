package models

import (
	"log"
	"time"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}


func (u *User) CreateUser() (err error) {
    cmd := `insert into users (
        uuid,
        name,
        email,
        password,
        created_at) values (?, ?, ?, ?, ?)`
    _, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())

    if err != nil {
        log.Fatal(err)
    }

    return err
}


func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return user, err
}


func (user *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, user.Name, user.Email, user.ID)

	if err != nil {
		log.Fatal(err)
	}

	return err
}


func (user *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, user.ID)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

