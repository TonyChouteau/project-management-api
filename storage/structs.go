package storage

// DATA

/*
ProjectCount struct
*/
type ProjectCount struct {
	NotStarted int `json:"notstarted"`
	Ongoing    int `json:"ongoing"`
	Closed     int `json:"closed"`
}

/*
CountByType struct
*/
type CountByType struct {
	Tags  map[string]ProjectCount `json:"tags"`
	Total int                     `json:"total"`
}

// PROJECTS

/*
Project struct
*/
type Project struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	SubTitle    string   `json:"subtitle"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	Type        int      `json:"type"`
	Solo        bool     `json:"solo"`
	Sources      string   `json:"sources"`
	Link        string   `json:"link"`
	Tags        []string `json:"tags"`
}

/*
ProjectList struct
*/
type ProjectList struct {
	Projects []Project `json:"project"`
	Count    int       `json:"count"`
}

/*
=====================
*/

/*
ImageByProject : list of name of each images in this project
*/
type ImageByProject []string

/*
ImageList : list of images sorted by project
*/
type ImageList map[int]ImageByProject
