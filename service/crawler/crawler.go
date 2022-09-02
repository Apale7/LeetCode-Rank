package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	db_model "github.com/Apale7/LeetCode-Rank/db/model"
	"github.com/Apale7/LeetCode-Rank/model"
	"github.com/Apale7/LeetCode-Rank/service/proxy"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	apiURL = "https://leetcode-cn.com/graphql"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	lang   map[string]string
	header map[string][]string
)

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
	header["user-agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/98.0.1108.56"}
	header["x-csrftoken"] = []string{"AFrbuAoCqSd8oN7A7AffwmmDgnYZ7V6uNolMLnJT5rcXEuiPlIGpNjzasr7eK85l"}
	header["x-definition-name"] = []string{"question"}
	header["x-operation-name"] = []string{"questionData"}
	header["x-timezone"] = []string{"Etc/Unknown"}
	// fmt.Println(header)
	lang = make(map[string]string)
	lang["A_0"] = "C++"
	lang["A_1"] = "Java"
	// lang["A_2"] = "Python2"
	// lang["A_3"] = "Python3"
	// lang["A_4"] = "C"
	// lang["A_5"] = "C#"
	// lang["A_6"] = "JavaScript"
	// lang["A_7"] = "Ruby"
	// lang["A_8"] = "Swift"
	lang["A_10"] = "Go"
	// lang["A_10"] = "Scala"
	lang["A_11"] = "Python3"
	// lang["A_12"] = "Rust"
	// lang["A_13"] = "PHP"
	lang["A_20"] = "TypeScript"
}

func unique(submits []model.RecentSubmissions) []model.RecentSubmissions {
	ans := []model.RecentSubmissions{}
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

func GetData(username string) []model.RecentSubmissions {
	// username := "apale"
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
	var data model.Info
	err = json.Unmarshal(body, &data)
	// fmt.Println(data)
	submmits := unique(data.Data.RecentSubmissions)
	return submmits
}

func GetDifficulty(title string) int {
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
	var data model.QuestionLevelInfo
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

func GetUserPublicProfile(username string) *model.AcData {
	fmt.Println(username)
	postData := model.PostData{
		OprationName: "userPublicProfile",
		Variables:    model.UserSlug{UserSlug: username},
		Query:        "query userPublicProfile($userSlug: String!) {userProfilePublicProfile(userSlug: $userSlug) {username submissionProgress {totalSubmissions waSubmissions acSubmissions reSubmissions otherSubmissions acTotal questionTotal __typename} __typename}}",
	}

	client := &http.Client{}
	pAddr, err := proxy.GetProxy()

	if err == nil && len(pAddr) > 0 {
		p, _ := url.Parse(pAddr)
		netTransport := &http.Transport{
			Proxy:                 http.ProxyURL(p),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		}
		client.Transport = netTransport
		log.Info("使用代理:", pAddr)
	}
	bytes, _ := json.Marshal(postData)

	res, err := client.Post(apiURL, "application/json", strings.NewReader(string(bytes)))
	if err != nil {
		log.Error(errors.WithStack(err))
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(errors.WithStack(err))
	}
	var data model.AcData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error(errors.WithStack(err))
		return nil
	}
	fmt.Printf("%+v", data)
	// submmits := unique(data.Data.RecentSubmissions)
	return &data
}

func GetUserQuestionProgress(username string) *db_model.Accepted {
	postData := model.PostData{
		OprationName: "userQuestionProgress",
		Variables:    model.UserSlug{UserSlug: username},
		Query:        "query userQuestionProgress($userSlug: String!) {\n  userProfileUserQuestionProgress(userSlug: $userSlug) {\n    numAcceptedQuestions {\n      difficulty\n      count\n      __typename\n    }\n    numFailedQuestions {\n      difficulty\n      count\n      __typename\n    }\n    numUntouchedQuestions {\n      difficulty\n      count\n      __typename\n    }\n    __typename\n  }\n}\n",
	}
	client := &http.Client{Timeout: time.Second * 10}

	bytes, _ := json.Marshal(postData)

	res, err := client.Post(apiURL, "application/json", strings.NewReader(string(bytes)))
	if err != nil {
		log.Error(errors.WithStack(err))
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(errors.WithStack(err))
	}
	data := model.QuestionInfo{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error(errors.WithStack(err))
		return nil
	}
	ret := db_model.NewAccepted()
	for _, info := range data.Data.UserProfileUserQuestionProgress.NumAcceptedQuestions {
		switch info.Difficulty {
		case "EASY":
			ret.Easy = int(info.Count)
		case "MEDIUM":
			ret.Medium = int(info.Count)
		case "HARD":
			ret.Hard = int(info.Count)
		}
	}
	return ret
}
