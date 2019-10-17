package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

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
GetProject : get a single project
*/
func GetProject(id int) Project {

	jsonFile, err := os.Open("data/projects.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	projects := []Project{}
	err = json.Unmarshal(byteValue, &projects)

	if id > 0 && id < len(projects)+1 {
		project := projects[id-1]
		return project
	}

	return Project{
		ID:          -1,
		Title:       "ERROR",
		SubTitle:    "No project with this ID",
		Description: "Id of project are between 1 and " + strconv.Itoa(len(projects)),
	}
}

/*
GetImageLists : get list of images sorted by project
*/
func GetImageLists() ImageList {

	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	imageList := map[int]ImageByProject{}
	for _, f := range files {

		images, err := ioutil.ReadDir("./images/" + f.Name())
		if err != nil {
			//log.Fatal(err)
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

/*
GetImageList : get a single dir list of images
*/
func GetImageList(id string) ImageByProject {

	if exists, _ := exists("./images/" + id); exists {

		images, err := ioutil.ReadDir("./images/" + id)
		if err != nil {
			log.Fatal(err)
		}

		imageByProject := ImageByProject{}

		for _, f := range images {
			imageByProject = append(imageByProject, f.Name())
		}

		return imageByProject
	}

	return ImageByProject{}
}
