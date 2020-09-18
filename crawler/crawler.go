package crawler

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var lang map[string]string
var header map[string][]string

func init() {
	header = make(map[string][]string)
	header["accept"] = []string{"*/*"}
	header["accept-encoding"] = []string{"gzip", "deflate", "br"}
	header["accept-language"] = []string{"zh-CN"}
	header["content-length"] = []string{"1262"}
	header["content-type"] = []string{"application/json"}
	header["origin"] = []string{"https://leetcode-cn.com"}
	header["referer"] = []string{"https://leetcode-cn.com/problems/two-sum/"}
	header["sec-fetch-dest"] = []string{"empty"}
	header["sec-fetch-mode"] = []string{"cors"}
	header["sec-fetch-site"] = []string{"same-origin"}
	header["user-agent"] = []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36"}
	header["x-csrftoken"] = []string{"AFrbuAoCqSd8oN7A7AffwmmDgnYZ7V6uNolMLnJT5rcXEuiPlIGpNjzasr7eK85l"}
	header["x-definition-name"] = []string{"question"}
	header["x-operation-name"] = []string{"questionData"}
	header["x-timezone"] = []string{"Etc/Unknown"}
	//fmt.Println(header)
	lang = make(map[string]string)
	lang["A_0"] = "C++"
	lang["A_1"] = "Java"
	//lang["A_2"] = "Python2"
	//lang["A_3"] = "Python3"
	//lang["A_4"] = "C"
	//lang["A_5"] = "C#"
	//lang["A_6"] = "JavaScript"
	//lang["A_7"] = "Ruby"
	//lang["A_8"] = "Swift"
	lang["A_10"] = "Go"
	//lang["A_10"] = "Scala"
	lang["A_11"] = "Python3"
	//lang["A_12"] = "Rust"
	//lang["A_13"] = "PHP"
	lang["A_20"] = "TypeScript"
}

func unique(submits []RecentSubmissions) []RecentSubmissions {
	ans := []RecentSubmissions{}
	st := make(map[string]bool)
	for _, v := range submits {
		if v.Status != "A_10" {
			continue
		}
		if !st[v.Question.QuestionFrontendID] {
			v.Lang = lang[v.Lang]
			v.Status = "accepted"
			ans = append(ans, v)
			st[v.Question.QuestionFrontendID] = true
		}
	}
	return ans
}

func GetData(username string) []RecentSubmissions {
	//username := "apale"
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
	submmits := unique(data.Data.RecentSubmissions)
	return submmits
}
func GetDifficulty(title string) (int){
	client := &http.Client{}
	body := fmt.Sprintf(`{"operationName": "questionData",
    "variables": {
       "titleSlug": "%s"
    },
    "query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {difficulty}\n}\n"}`, title)
	req, _ := http.NewRequest("POST", "https://leetcode-cn.com/graphql/", strings.NewReader(body))
	req.Header = header
	res, _ := client.Do(req)
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(errors.WithStack(err))
	}
	var data QuestionLevelInfo
	err = json.Unmarshal(resBody, &data)
	switch data.Data.Question.Difficulty {
	case "Easy":
		return 0
	case "Medium":
		return 1
	default:
		return 2
	}
}

type QuestionLevelInfo struct {
	Data QuestionLevelData `json:"data"`
}

type QuestionLevel struct {
	Difficulty string `json:"difficulty"`
}
type QuestionLevelData struct {
	Question QuestionLevel `json:"question"`
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
	SubmitTime int64    `json:"submitTime"`
	Typename   string   `json:"__typename"`
}
type Data struct {
	RecentSubmissions []RecentSubmissions `json:"recentSubmissions"`
}
