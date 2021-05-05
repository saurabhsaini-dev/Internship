package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ListUsers struct {
	Results []User `json:"results"`
}

type User struct {
	Name          string `json:"name"`
	FullName      string `json:"fullname"`
	Email         string `json:"eamil"`
	AvatarUrl     string `json:"avatarUrl"`
	DirectoryName string `json:"directoryName"`
	IsActive      bool   `json:"isActive"`
	Editable      bool   `json:"editable"`
}

func getUsers() {
	httpClient := &http.Client{}
	url := "http://localhost:8085/rest/api/latest/admin/users/"
	httpMethod := http.MethodGet
	token := "MDM2NTg2ODM0NzUyOhbUXhoq8OZLU6elRsJ+ym9twxPb"

	var err error
	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("error")
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}

	var response []byte
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var apiResponse ListUsers
	err = json.Unmarshal(response, &apiResponse)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(apiResponse)

}

func main() {
	getUsers()
}
