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
	Type map[string]ProjectCount `json:"type"`
}

// PROJECTS

/*
Project struct
*/
type Project struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

/*
ProjectList struct
*/
type ProjectList struct {
	Project Project `json:"project"`
	Count   int     `json:"count"`
}
