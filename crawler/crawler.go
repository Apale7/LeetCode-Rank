package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)



var json = jsoniter.ConfigCompatibleWithStandardLibrary

func getData() {
	username := "apale"
	url := "https://leetcode-cn.com/graphql?oprationName=recentSubmissions&variables={%22userSlug%22:%22" + username + "%22}&query=query%20recentSubmissions($userSlug:%20String!){recentSubmissions(userSlug:%20$userSlug){status%20lang%20question{questionFrontendId%20title%20translatedTitle%20titleSlug%20__typename}submitTime%20__typename}}"
	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		log.Error(errors.WithStack(err))
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(errors.WithStack(err))
	}
	var data Info
	err = json.Unmarshal(body, &data)
	//fmt.Println(submittions)
	for i, submmit := range data.Data.RecentSubmissions {
		fmt.Println(i, submmit)
	}
}

func main() {
	getData()
}


type Info struct {
	Data Data `json:"data"`
}
type Question struct {
	QuestionFrontendID string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TranslatedTitle    string `json:"translatedTitle"`
	TitleSlug          string `json:"titleSlug"`
	Typename           string `json:"__typename"`
}
type RecentSubmissions struct {
	Status     string   `json:"status"`
	Lang       string   `json:"lang"`
	Question   Question `json:"question"`
	SubmitTime int      `json:"submitTime"`
	Typename   string   `json:"__typename"`
}
type Data struct {
	RecentSubmissions []RecentSubmissions `json:"recentSubmissions"`
}