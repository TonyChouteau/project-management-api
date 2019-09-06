package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

/*
CountProjects function
*/
func CountProjects(c *gin.Context) CountByType {

	msg := c.DefaultQuery("firstname", "false")
	fmt.Println(msg)

	jsonFile, err := os.Open("data/projects.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	projects := []Project{}
	err = json.Unmarshal(byteValue, &projects)

	typesMap := map[string]ProjectCount{}

	for _, project := range projects {
		for _, tag := range project.Tags {
			var pc ProjectCount
			if _, ok := typesMap[tag]; ok {
				pc = typesMap[tag]
			} else {
				pc = ProjectCount{0, 0, 0}
			}
			switch project.Status {
			case 0:
				pc.NotStarted++
			case 1:
				pc.Ongoing++
			case 2:
				pc.Closed++
			}
			typesMap[tag] = pc
		}
	}

	countByType := CountByType{}

	countByType.Tags = typesMap
	countByType.Total = len(projects)

	return countByType
}

/*
GetProjects function
*/
func GetProjects(c *gin.Context) ProjectList {

	jsonFile, err := os.Open("data/projects.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	projects := []Project{}
	err = json.Unmarshal(byteValue, &projects)

	return ProjectList{
		projects,
		len(projects),
	}
}
