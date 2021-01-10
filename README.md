# LeetCode-Rank
配置LeetCode的username列表后，定期爬取用户提交记录，存入MySQL和redis
## 运行
- go mod tidy
- go build -o rank
- ./rank

## 配置
+ config文件夹下
  + dbconf.yml中配置MySQL
  + userlist.yml配置要爬取的用户列表
  + user_map.yml配置username: nickname的键值对，接口返回的name是nickname
  + localhost:6379必须有redis，否则启动时panic

## 接口
+ localhost:4398

##目录
<ul>
<li>├── config 配置文件</li> 
<li>├── crawler爬虫</li> 
<li>├── db MySQL、Redis</li>
<li>├── go.mod</li>
<li>├── go.sum</li>
<li>├── handler 接口实现</li>
<li>├── main.go</li>
<li>├── model 结构体</li>
<li>├── router.go 接口路由</li>
└── utils
</ul>
