package https

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
1) Getters & Setters -
	a - get ongoing/closed projects count (by type) : Web, Processing3, Android, Unity, Other (Java, AI, )
	b - get projects : Get Name, Image, Description, Tags
	c - get projects with search engine : By Name, Tags, Ongoing/Closed
2) Post -
	- post project : Give Name, Image, Description and Tags
3) Store projects -
	- store project in a directory

Annexes -
	Status : 0 not started yet, 1 closed, 2 ongoing
*/

/*
Count Getter : See "1)a"
*/
func counter(c *gin.Context) {

	c.JSON(200, gin.H{
		"type": typeParams,
	})
}

/*
Serve function
*/
func Serve() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/count", counter)
	//r.StaticFile("/image", "./images/951546.jpg")

	err := http.ListenAndServe(":8081", r)
	//err := http.ListenAndServeTLS(":8081", "/etc/letsencrypt/live/www.domain.com/fullchain.pem", "/etc/letsencrypt/live/www.domain.com/privkey.pem", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
