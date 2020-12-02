package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main()  {
	id := uint(1)
	user := UserServer{}
	user.findUserById(id)
	fmt.Println(user)

}

//Service层
type UserServer struct {
	UserDao *UserDao
}

func (u *UserServer) findUserById(id uint) *sql.Result {
	user, err := u.UserDao.findUserById(id)
	if err != nil {
		fmt.Printf("error info %v", err)
		return nil
	}
	return  user
}

//Dao层

type UserDao struct {
	Db *sql.DB

}

func (u *UserDao) findUserById(id uint) (*sql.Result, error) {
	user, err := u.Db.Exec("select * from user where id = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("findUserById error:" + sql.ErrNoRows.Error())
	}
	if err != nil {
		return nil, errors.Wrap(err,"findUserById error ...")
	}
	return &user, nil
}