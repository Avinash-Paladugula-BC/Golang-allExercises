package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	env "gorm/env_variables"
)

func create(db *gorm.DB) {
	var newPost Post
	newPost = Post{
		Title:   "new title",
		Content: "new content",
	}
	result := db.Create(&newPost)
	if result.Error != nil {
		fmt.Errorf("unable to retrieve the records : ", result.Error)
	}
	read(db)
}

func read(db *gorm.DB) {
	var posts []Post
	result := db.Find(&posts)
	if result.Error != nil {
		fmt.Errorf("unable to retrieve the records : ", result.Error)
	}
	for _, post := range posts {
		fmt.Printf("ID: %d \nPost Title: %s \nPost Content: %s \n\n", post.ID, post.Title, post.Content)
	}
}

func update(db *gorm.DB) {
	result := db.Model(&Post{}).Where("id = ?", 2).Updates(Post{Title: "New title", Content: "new content"})
	if result.Error != nil {
		fmt.Errorf("Error occured while updating")
	}
}

func delete(db *gorm.DB) {
	result := db.Delete(&Post{}, 1)
	if result.Error != nil {
		fmt.Errorf("Error occured while updating")
	}
}

func search(db *gorm.DB, id uint) {
	var posts []Post
	db.Where("id = ?", id).Find(&posts)
	for _, post := range posts {
		fmt.Printf("ID: %d \nPost Title: %s \nPost Content: %s \n\n", post.ID, post.Title, post.Content)
	}

}

func main() {

	dbName := os.Args[1]
	fmt.Println(dbName)

	sslmode := "disable"
	creds := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", env.DbUserName, env.DbPassword, env.DbName, sslmode)
	db, err := gorm.Open("postgres", creds)

	if err != nil {
		fmt.Errorf("Error occured : ", err)
	}
	db.DropTable(&Post{})
	db.CreateTable(&Post{})
	for _, post := range defaultPosts {
		db.Create(&post)
	}
	fmt.Println("creation.........")
	create(db)
	fmt.Println("reading.........")
	read(db)
	fmt.Println("update.........")
	update(db)
	read(db)
	fmt.Println("delete.........")
	delete(db)
	read(db)
	fmt.Println("search.........")
	search(db, 3)

}

type Post struct {
	gorm.Model
	Title   string `gorm: "not null; uniqueIndex"`
	Content string `gorm: "not null; unique"`
}

var defaultPosts []Post = []Post{
	Post{Title: "Fitness", Content: "Focus on fitness"},
	Post{Title: "Science", Content: "Learning related to Sciece"},
	Post{Title: "Sports", Content: "Sports updates"},
	Post{Title: "Business", Content: "Business updates"},
	Post{Title: "Technical", Content: "Technical knowledge"},
}
