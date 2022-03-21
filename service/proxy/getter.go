package proxy

import "github.com/Apale7/lazy_proxy/proxy_getter"

var getter proxy_getter.AutoProxyGetter

func init() {
	return
	getter = &proxy_getter.DefaultAutoProxyGetter{
		ProxyGetter: &proxy_getter.DefaultProxyGetter{},
		Crawler:     &proxy_getter.CrawlerIP3366{},
	}
	getter = proxy_getter.WrapWithTimeDecorator(getter, 300)
//	getter = proxy_getter.WrapWithThresholdDecorator(getter, 10)
	getter.PushProxy(getter.CrawlProxy()...)

	if getter.LenOfProxies() == 0 {
		panic("no proxy")
	}
}

func GetProxy() (string, error) {
	return "", nil
//	return getter.GetProxy()
}
