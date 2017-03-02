package models

import (
	"time"
)

type ShopCar struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	BookId    int       `json:"bookId"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

func (this *ShopCar) Add() error {
	return insert(this)
}

func (this *ShopCar) Update() error {
	return update(this, this.Id)
}

func (this *ShopCar) Del() error {
	return delete(this, this.Id)
}

func MyShopcarBookIds(uId int) (r []int, err error) {
	var shopcars []*ShopCar
	err = engine.Where("user_id=?", uId).Find(&shopcars)
	if err != nil {
		return
	}
	for _, item := range shopcars {
		r = append(r, item.BookId)
	}
	return
}
