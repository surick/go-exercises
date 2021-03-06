package main

import (
	"log"

	"github.com/surick/go-exercises/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
	db.CreateTable(model.User{}, model.Post{})

	model.AddUser("surick", "123456", "jk103@qq.com")
	model.AddUser("jin", "123456", "root@jinaiya.com")

	u1, _ := model.GetUserByUsername("surick")
	u1.CreatePost("Beautiful day in Portland!")
	model.UpdateAboutMe(u1.Username, `I'm the author of Go-exercises Tutorial you are reading now!`)

	u2, _ := model.GetUserByUsername("jin")
	u2.CreatePost("The Avengers movie was so cool!")
	u2.CreatePost("Sun shine is beautiful")

	u1.Follow(u2.Username)
}