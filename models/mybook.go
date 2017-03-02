package models

import (
	"time"
)

type MyBook struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	UserId    int       `json:"userId"`
	BookId    int       `json:"bookId"`
	CreatedAt time.Time `xorm:"created"`
}

func (this *MyBook) Add() error {
	return insert(this)
}

func (this *MyBook) Update() error {
	return update(this, this.Id)
}

func (this *MyBook) Del() error {
	return delete(this, this.Id)
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

func MyBookIds(uId int) (r []int, err error) {
	var books []*MyBook
	err = engine.Where("user_id=?", uId).Find(&books)
	if err != nil {
		return
	}
	for _, item := range books {
		r = append(r, item.BookId)
	}
	return
}
