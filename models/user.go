package models

import (
	"books/utils"
	"errors"
	"fmt"
	"time"
)

type User struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	UserName  string    `json:"username"`
	PassWord  string    `json:"password,omitempty"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Pic       string    `json:"pic"`
	Money     float64   `json:"money"`
	Token     string    `json:"token"`
	CreatedAt time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

func (this *User) Add() error {
	this.Token = utils.EncodePasswd(fmt.Sprintf("%s_%d", this.UserName, time.Now().Unix()))
	return insert(this)
}

func (this *User) Update() error {
	_, err := engine.Where("id=?", this.Id).Update(this)
	return err
}

func (this *User) Del() error {
	return delete(this)
}

func (this *User) Get() error {
	_, err := engine.Get(this)
	if err != nil {
		return err
	}
	this.PassWord = ""
	return nil
}

func GetUserByToken(token string) (user *User, err error) {
	user = new(User)
	has, err := engine.Where("token=?", token).Get(user)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	if !has {
		return nil, errors.New("not found user.")
	}
	user.PassWord = ""
	return
}

func Login(userName, email, pwd string) (user *User, err error) {
	user = new(User)
	has, err := engine.Where("(user_name=? OR email=?) AND pass_word=?", userName, email, pwd).Get(user)
	if err != nil {
		return
	}
	if !has {
		return nil, errors.New("not found user.")
	}
	user.Token = utils.EncodePasswd(fmt.Sprintf("%s_%d", userName, time.Now().Unix()))
	err = user.Update()
	if err != nil {
		return
	}
	user.PassWord = ""
	return
}

func GetUserInId(ids []int) (r []*User, err error) {
	err = engine.In("id", ids).Find(&r)
	return
}
