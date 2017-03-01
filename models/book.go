package models

import (
	"errors"
	"time"
)

type Book struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	Name      string    `json:"name"`
	Autor     string    `json:"autor"`
	Desc      string    `json:"desc"`
	Pic       string    `json:"pic"`
	Money     float64   `json:"money"`
	Category  int       `json:"category"`
	CreatedAt time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

func (this *Book) Add() error {
	return insert(this)
}

func (this *Book) Update() error {
	_, err := engine.Where("id=?", this.Id).Update(this)
	return err
}

func (this *Book) Del() error {
	return delete(this)
}

func (this *Book) Get() error {
	has, err := engine.Get(this)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("not found book.")
	}
	return nil
}

func BookList() (r []*Book, err error) {
	err = engine.Find(&r)
	return
}

func BookListByCategory(t int) (r []*Book, err error) {
	err = engine.Where("category=?", t).Find(&r)
	return
}

func BookInId(ids []int) (r []*Book, err error) {
	err = engine.In("id", ids).Find(&r)
	return
}
