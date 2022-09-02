package proxy

import (
	"testing"
)

func Test_jiangxiliCralwer_CrawlProxy(t *testing.T) {
	c := &jiangxiliCralwer{}
	p := c.CrawlProxy()
	t.Log(p)
}
