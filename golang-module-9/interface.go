package main

import "github.com/jinzhu/gorm"

type Store interface {
	create(db *gorm.DB)
	read(db *gorm.DB)
	update(db *gorm.DB, book_id uint)
	delete(db *gorm.DB, book_id uint)
	searchByAuthor(db *gorm.DB, author string)
	searchByTitle(db *gorm.DB, title string)
	search(db *gorm.DB, id uint)
}
