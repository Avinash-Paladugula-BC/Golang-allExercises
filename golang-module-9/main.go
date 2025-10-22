package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

)

func create(db *gorm.DB) {
	// var newPost Post
	newPost := Post{
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
		fmt.Errorf("Error occured while deleting")
	}
}

func searchByAuthor(db *gorm.DB, author string) {
	var post Post
	db.Where("author = ?", author).First(&post)
	fmt.Printf("ID: %d \nPost Title: %s \nPost Content: %s \n\n", post.ID, post.Title, post.Content)

}

func searchByTitle(db *gorm.DB, title string) {
	var post Post
	db.Where("title = ?", title).First(&post)
	fmt.Printf("ID: %d \nPost Title: %s \nPost Content: %s \n\n", post.ID, post.Title, post.Content)

}

func search(db *gorm.DB, id uint) {
	var post Post
	db.Where("id = ?", id).First(&post)
	fmt.Printf("ID: %d \nPost Title: %s \nPost Content: %s \n\n", post.ID, post.Title, post.Content)
	// for _, post := range posts {
	// 	fmt.Printf("ID: %d \nPost Title: %s \nPost Content: %s \n\n", post.ID, post.Title, post.Content)
	// }

}

func main() {
	// Run the below lines and set the environment variables
	// export DB_NAME="mydb"
	// export SSLMODE="disable"
	// export USER="postgres"
	// export PASSWORD="Avi@2004"

	
	creds := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSLMODE"))
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
	fmt.Println("search by author name.........")
	searchByAuthor(db, "Vinay")
	fmt.Println("search by title.........")
	searchByAuthor(db, "Science")

}

type Post struct {
	gorm.Model
	Title   string `gorm: "not null; uniqueIndex"`
	Content string `gorm: "not null; unique"`
	Author string 
}

var defaultPosts []Post = []Post{
	Post{Title: "Fitness", Content: "Focus on fitness" , Author: "Vinay"},
	Post{Title: "Science", Content: "Learning related to Sciece" , Author: "Rasagnya"},
	Post{Title: "Sports", Content: "Sports updates" , Author: "Akhila"},
	Post{Title: "Business", Content: "Business updates" , Author: "Vinay"},
	Post{Title: "Technical", Content: "Technical knowledge" , Author: "Akhila"},
}
