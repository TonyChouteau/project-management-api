package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

/*
CountProjects function
*/
func CountProjects() CountByType {

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
func GetProjects() ProjectList {

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

/*
GetImageList : get list of images sorted by project
*/
func GetImageList() ImageList {

	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	imageList := map[int]ImageByProject{}
	for _, f := range files {

		images, err := ioutil.ReadDir("./images/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		projectImages := ImageByProject{}
		for _, f := range images {
			projectImages = append(projectImages, f.Name())
		}

		key, err := strconv.Atoi(f.Name())
		if err != nil {
			log.Fatal(err)
		}
		imageList[key] = projectImages
	}

	return imageList
}
