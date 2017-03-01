package models

import (
	"time"
)

type MyBook struct {
	Id        int `xorm:"pk autoincr" json:"id"`
	UserId    int
	BookId    int
	CreatedAt time.Time `xorm:"created"`
}

func (this *MyBook) Add() error {
	return insert(this)
}

func (this *MyBook) Update() error {
	return update(this)
}

func (this *MyBook) Del() error {
	return delete(this)
}

func BuyBook(bId, uId int) (book *Book, err error) {
	myBook := &MyBook{UserId: uId, BookId: bId}
	if err = myBook.Add(); err != nil {
		return
	}
	sc := &ShopCar{BookId: bId, UserId: uId}
	if err = sc.Del(); err != nil {
		return
	}
	book = &Book{Id: bId}
	if err = book.Get(); err != nil {
		return nil, err
	}
	return
}
