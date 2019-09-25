package https

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TonyChouteau/project-management-api/storage"
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
	Status : 0 not started yet, 1 ongoing, 2 closed
*/

/*
Count Getter : See "1)a"
*/
func counter(c *gin.Context) {

	countByType := storage.CountProjects(c)

	c.JSON(200, countByType)
}

func projects(c *gin.Context) {

	projects := storage.GetProjects(c)

	c.JSON(200, projects)
}

func project(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	project := storage.GetProject(c, id)

	c.JSON(200, project)
}

/*
Serve function
*/
func Serve() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/count", counter)
	r.GET("/projects", projects)
	r.GET("/project/:id", project)
	//r.StaticFile("/image", "./images/951546.jpg")

	err := http.ListenAndServe(":8081", r)
	//err := http.ListenAndServeTLS(":8081", "/etc/letsencrypt/live/www.domain.com/fullchain.pem", "/etc/letsencrypt/live/www.domain.com/privkey.pem", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
