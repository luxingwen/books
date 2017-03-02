package models

import (
	"time"
)

type Community struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	Content   string    `json:"content"`
	BookId    int       `json:"bookId"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `xorm:"created"`
}

func (this *Community) Add() error {
	return insert(this)
}

func (this *Community) Update() error {
	return update(this, this.Id)
}

func (this *Community) Del() error {
	return delete(this, this.Id)
}

func CommunityListByBook(bookId int) (r []*Community, err error) {
	err = engine.Where("book_id=?", bookId).Find(&r)
	return
}
