package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitProfileData struct {
	UserName    string `json:"login"`
	AvatarUrl   string `json:"avatar_url"`
	ProfileUrl  string `json:"url"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Email       string `json:"email"`
	Followers   int64  `json:"followers"`
	Following   int64  `json:"following"`
	Created     string `json:"created_at"`
	LastUpdated string `json:"updated_at"`
}

func main() {
	fmt.Println("Github Stats App")

	result, err := http.Get("https://api.github.com/users/niteshkamble")

	if err != nil {
		println("Failed HTTP Request : ", err)
	}

	defer result.Body.Close()

	if result.StatusCode != 200 {
		panic("Github Api Failed.")
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	var gitProfileData GitProfileData
	err = json.Unmarshal(body, &gitProfileData)

	// fmt.Println(gitProfileData)

	if err != nil {
		panic(err)
	}

	username, name, comapany, location, followers, following, email :=
		gitProfileData.UserName, gitProfileData.Name, gitProfileData.Company, gitProfileData.Location, gitProfileData.Followers, gitProfileData.Following, gitProfileData.Email

	fmt.Printf(
		"\nUsername : %s \nName : %s \nCompany : %s \nLocation : %s \nFollowers : %d \nFollowing : %d \nEmail : %s\n",
		username, name, comapany, location, followers, following, email)

}
