package main

import (
	"log"
	"net/http"
	"github.com/prathyushnallamothu/usernamequery/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.GET("/search/:username", searchhandler)
	router.GET("/test1", test1handler)
	router.GET("/test2",test2handler)
	router.Run()

}

func searchhandler(c *gin.Context) {
	var Username, DBname string
	username := c.Param("username")
	db := database.Connect("test")
	defer db.Close()
	result, err := db.Query("select username,dbname from users where username=?", username)
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		err = result.Scan(&Username, &DBname)
		if err != nil {
			log.Fatal(err)
		}
	}
	if Username == "" && DBname == "" {
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "no results",
			})
	}
	if DBname == "test1" {
		c.Redirect(307, "/test1?username="+Username)
	}
	if DBname == "test2" {
		c.Redirect(307, "/test2?username="+Username)
	}
}
func test1handler(c *gin.Context){
	var Username,Email,Age,Location string
	username:=c.Query("username")
	db := database.Connect("test1")
	defer db.Close()
	result, err := db.Query("select username,email,age,location from users where username=?", username)
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		err = result.Scan(&Username, &Email,&Age,&Location)
		if err != nil {
			log.Fatal(err)
		}
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"username":Username,
			"email":Email,
			"age":Age,
			"location":Location,
		})	
}
func test2handler(c *gin.Context){
	var Username,Email,Age,Location string
	username:=c.Query("username")
	db := database.Connect("test2")
	defer db.Close()
	result, err := db.Query("select username,email,age,location from users where username=?", username)
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		err = result.Scan(&Username, &Email,&Age,&Location)
		if err != nil {
			log.Fatal(err)
		}
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"username":Username,
			"email":Email,
			"age":Age,
			"location":Location,
		})	
}