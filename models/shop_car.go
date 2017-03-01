package models

import (
	"time"
)

type ShopCar struct {
	Id        int `xorm:"pk autoincr" json:"id"`
	BookId    int
	UserId    int
	CreatedAt time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

func (this *ShopCar) Add() error {
	return insert(this)
}

func (this *ShopCar) Update() error {
	return update(this)
}

func (this *ShopCar) Del() error {
	return delete(this)
}
