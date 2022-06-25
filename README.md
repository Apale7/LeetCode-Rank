# LeetCode-Rank

定期爬取用户提交记录，存入mongodb
list接口有单机缓存

## 运行

- go mod tidy
- go build -o rank
- ./rank

## 配置

+ config/conf.yml
  + mongo_uri: mongodb:xxxxxxxxxxxxxxxx

## 接口

+ localhost:4398
