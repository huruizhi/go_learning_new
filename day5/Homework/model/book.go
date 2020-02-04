package model

import (
	"errors"
	"sync"
)

var ErrStockNotEnough = errors.New("stock is not enough")

type Book struct {
	Name       string
	Total      uint
	Author     string
	CreateTime string
	lock       sync.Mutex
}

func CreateBook(name string, total uint, author string, createTime string) (b *Book) {
	b = &Book{
		Name:       name,
		Author:     author,
		Total:      total,
		CreateTime: createTime,
	}
	return b
}

func (b *Book) CanBorrow(c uint) bool {
	return b.Total >= c
}

func (b *Book) Borrow(c uint) (err error) {
	b.lock.Lock()
	if b.CanBorrow(c) {
		b.Total = -c
	} else {
		err = ErrStockNotEnough
		return err
	}
	b.lock.Unlock()
	return
}

func (b *Book) Back(c uint) {
	b.lock.Lock()
	b.Total += c
	b.lock.Unlock()
}
