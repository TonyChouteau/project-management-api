package main

import "github.com/TonyChouteau/project-managment-api/https"

/*
1) Getters & Setters -
	- get ongoing/closed projects count (by type) : Web, Processing3, Android, Unity, Autre (Java, AI, )
	- get projects : Get Name, Image, Description, Tags
	- get projects with search engine : By Name, Tags, Ongoing/Closed
2) Post -
	- post project : Give Name, Image, Description and Tags
3) Store projects -
	- store project in a directory
*/

func main() {
	https.Serve()
}
