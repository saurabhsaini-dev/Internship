package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ListUsers struct {
	PageCount     int    `json:"page_count"`
	PageNumber    int    `json:"page_number"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	NextPageToken string `json:"next_page_token"`
	Users         []User `json:"users"`
}

type User struct {
	Id            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Type          int    `json:"type"`
	PMI           int    `json:"pmi"`
	Verified      int    `json:"verified"`
	CreatedAt     string `json:"created_at"`
	LastLoginTime string `json:"last_login_time"`
	PicUrl        string `json:"pic_url"`
	Status        string `json:"status"`
	RoleId        string `json:"role_id"`
}

func handleRequest() (response []byte, err error) {
	httpClient := &http.Client{}
	url := "https://api.zoom.us/v2/users/"
	httpMethod := http.MethodGet
	authToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTg4MTA2NzQsImlhdCI6MTYxODgwNTI3Nn0.Oty8DXzssVaCcP2jdqjFTH2hQxst5CsBEyp95bkCsKI"

	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return

}

func main() {
	res, err := handleRequest()
	if err != nil {
		fmt.Println("error")
	}

	var apiResponse ListUsers
	err = json.Unmarshal(res, &apiResponse)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(apiResponse)
}
