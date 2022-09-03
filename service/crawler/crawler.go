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
	"github.com/bytedance/sonic"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	apiURL = "https://leetcode-cn.com/graphql"
)

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
	bytes, _ := sonic.Marshal(postData)

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
	err = sonic.Unmarshal(body, &data)
	if err != nil {
		log.Error(errors.WithStack(err))
		return nil
	}
	fmt.Printf("%+v", data)
	// submmits := unique(data.Data.RecentSubmissions)
	return &data
}

func GetUserQuestionProgress(username string) *db_model.Accepted {
	url := "https://leetcode.cn/graphql/"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"query":"\n    query userQuestionProgress($userSlug: String!) {\n  userProfileUserQuestionProgress(userSlug: $userSlug) {\n    numAcceptedQuestions {\n      difficulty\n      count\n    }\n    numFailedQuestions {\n      difficulty\n      count\n    }\n    numUntouchedQuestions {\n      difficulty\n      count\n    }\n  }\n}\n    ","variables":{"userSlug":"%s"}}`, username))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Error(errors.WithStack(err))
		return nil
	}
	req.Header.Add("authority", "leetcode.cn")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("authorization", "")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "gr_user_id=aff0091b-7014-4a5b-9fac-28d37be10a75; _bl_uid=vXl3y37q40wlwR8ja814y5FtFLvR; a2873925c34ecbd2_gr_last_sent_cs1=apale; aliyungf_tc=ab3acadfaafb0bfdd49a76930438b947270fd1210c029645b160f3da0dc5a5fd; NEW_PROBLEMLIST_PAGE=1; a2873925c34ecbd2_gr_session_id=4c8a52fa-0d4f-4e68-9eb3-2d58dfdefc7b; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=4c8a52fa-0d4f-4e68-9eb3-2d58dfdefc7b; a2873925c34ecbd2_gr_session_id_4c8a52fa-0d4f-4e68-9eb3-2d58dfdefc7b=true; csrftoken=4SZIVBNMtbjWTXEOZXd2h8NNUqdkCIRbm62TA6vJkJe7EVbz3Sj4yjxzqldhuYQi; messages=\"[[\\\"__json_message\\\"\\0540\\05425\\054\\\"\\\\u60a8\\\\u5df2\\\\u7ecf\\\\u767b\\\\u51fa\\\"]]:1oUQ9l:DmtM_12_NGwKhhecx9R5C7Jv2krlscbSzfP2VVJj2Lg\"; a2873925c34ecbd2_gr_cs1=apale")
	req.Header.Add("origin", "https://leetcode.cn")
	req.Header.Add("referer", "https://leetcode.cn/u/apale/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Microsoft Edge\";v=\"104\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.70")
	req.Header.Add("x-csrftoken", "4SZIVBNMtbjWTXEOZXd2h8NNUqdkCIRbm62TA6vJkJe7EVbz3Sj4yjxzqldhuYQi")

	res, err := client.Do(req)
	if err != nil {
		log.Error(errors.WithStack(err))
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(errors.WithStack(err))
		return nil
	}

	data := model.QuestionInfo{}
	err = sonic.Unmarshal(body, &data)
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
