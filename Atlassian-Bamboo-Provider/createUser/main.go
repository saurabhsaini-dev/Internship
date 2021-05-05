package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUser struct {
	Name            string `json:"name"`
	FullName        string `json:"fullname"`
	Email           string `json:"eamil"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func createUser() {

	createUser := CreateUser{
		Name:            "perseverence",
		FullName:        "mars rover",
		Email:           "saurabhsaini@gmial.com",
		Password:        "xyz@123",
		PasswordConfirm: "xyz@123",
	}

	var err error

	var reqBody []byte
	reqBody, err = json.Marshal(createUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpClient := &http.Client{}
	url := "http://localhost:8085/rest/api/latest/admin/users/"
	httpMethod := http.MethodPost
	token := "MDM2NTg2ODM0NzUyOhbUXhoq8OZLU6elRsJ+ym9twxPb"

	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(reqBody))
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

	if resp.StatusCode != 204 {
		fmt.Println("error")
		fmt.Println(resp.StatusCode)
		return
	}

	if resp.StatusCode == 204 {
		fmt.Println("ok")
		return
	}

	/*var response []byte
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

	fmt.Println(apiResponse)*/

}

func main() {
	createUser()
}

