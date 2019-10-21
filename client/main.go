package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Book struct {
	Title string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	var tmp string
	userName := "gest user"
	var userID int64
	var book []Book
	var bookID int64
	var lendOrBorrow string
	var sw bool
	var err error
	var title string

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.POST("/begin", func(ctx *gin.Context) {
		userName = ctx.PostForm("userName")
		userID, err = strconv.ParseInt(userName, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		lendOrBorrow = ctx.PostForm("sw")
		sw = lendOrBorrow == "lend"
		ctx.Redirect(302, "/rental")
	})

	router.GET("/rental", func(ctx *gin.Context) {
		fmt.Println(book)
		ctx.HTML(200, "rental.html", gin.H{
			"userName": userName,
			"book":     book,
		})
	})

	router.POST("/addbook", func(ctx *gin.Context) {
		tmp = ctx.PostForm("bookName")
		bookID, err = strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			book = append(book, Book{Title: "error"})
			ctx.Redirect(302, "/rental")
		}
		title, err = client(userID, bookID, sw)
		if err != nil {
			if len(book) == 0 {
				book = append(book, Book{Title: "error"})
			} else if book[len(book)-1].Title == "error" {
				book[len(book)-1].Title = "error"
			} else {
				book = append(book, Book{Title: "error"})
			}
		} else {
			if len(book) == 0 {
				book = append(book, Book{Title: title})
			} else if book[len(book)-1].Title == "error" {
				book[len(book)-1].Title = title
			} else {
				book = append(book, Book{Title: title})
			}
		}
		fmt.Println(title)
		ctx.Redirect(302, "/rental")
	})

	router.Run()
}
