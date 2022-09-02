package proxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type jiangxiliCralwer struct {
}

func (c *jiangxiliCralwer) CrawlProxy() []string {
	resp, err := http.Get("https://ip.jiangxianli.com/api/proxy_ips")
	if err != nil {
		logrus.Errorf("jiangxiliCralwer CrawlProxy error: %v", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("jiangxiliCralwer CrawlProxy error: %v", err)
		return nil
	}
	var res Resp

	if err := json.Unmarshal(body, &res); err != nil {
		logrus.Errorf("jiangxiliCralwer CrawlProxy error: %v", err)
		return nil
	}
	if res.Code != 0 {
		logrus.Errorf("jiangxiliCralwer CrawlProxy error: %v", res.Msg)
		return nil
	}
	var proxyList []string
	for _, a := range res.Data.Data {
		proxyList = append(proxyList, a.IP+":"+a.Port)
	}
	return proxyList
}

type Resp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	CurrentPage  int64       `json:"current_page"`
	Data         []Datum     `json:"data"`
	FirstPageURL string      `json:"first_page_url"`
	From         int64       `json:"from"`
	LastPage     int64       `json:"last_page"`
	LastPageURL  string      `json:"last_page_url"`
	NextPageURL  interface{} `json:"next_page_url"`
	Path         string      `json:"path"`
	PerPage      int64       `json:"per_page"`
	PrevPageURL  interface{} `json:"prev_page_url"`
	To           int64       `json:"to"`
	Total        int64       `json:"total"`
}

type Datum struct {
	UniqueID    string `json:"unique_id"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
	Country     string `json:"country"`
	IPAddress   string `json:"ip_address"`
	Anonymity   int64  `json:"anonymity"`
	Protocol    string `json:"protocol"`
	ISP         string `json:"isp"`
	Speed       int64  `json:"speed"`
	ValidatedAt string `json:"validated_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
