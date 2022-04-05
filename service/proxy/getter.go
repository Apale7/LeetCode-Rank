package proxy

import (
	"fmt"

	"github.com/Apale7/lazy_proxy/proxy_pool"
	"github.com/Apale7/lazy_proxy/proxy_pool/proxy_crawler"
)

var getter proxy_pool.AutoProxyGetter

func init() {
	fmt.Println("proxy init")
	getter = &proxy_pool.DefaultAutoProxyGetter{
		ProxyPool: &proxy_pool.DefaultProxyPool{},
		Crawler:   &proxy_crawler.CrawlerKuaidaili{},
	}
	getter = proxy_pool.WrapWithTimeDecorator(getter, 300)
	getter = proxy_pool.WrapWithThresholdDecorator(getter, 25)
	getter.PushProxy(getter.CrawlProxy()...)

	if getter.LenOfProxies() == 0 {
		panic("no proxy")
	}
	fmt.Println("proxy init done")
}

func GetProxy() (string, error) {
	return getter.GetProxy()
}
