# 启动

```go
go run .
```

# API

## 获取所有

```sh
curl http://localhost:8080/albums
curl http://localhost:8080/albums -H "Content-Type: application/json" -X "GET"
curl http://localhost:8080/albums --header "Content-Type: application/json" --request "GET"
```

## 获取单个

```sh
curl http://localhost:8080/albums/2
```

## 新增单条

```sh
# PowerShell命令行不换行方式新增单条
curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{\"id\": \"4\",\"title\": \"The Modern Sound of Betty Carter\",\"artist\": \"Betty Carter\",\"price\": 49.99}'

# PowerShell命令行换行方式新增单条
curl http://localhost:8080/albums `
    --include `
    --header "Content-Type: application/json" `
    --request "POST" `
    --data '[{\"id\": \"10\",\"title\": \"The Modern Sound of Betty Carter\",\"artist\": \"Betty Carter\",\"price\": 49.99,\"id\": \"11\",\"title\": \"The Modern Sound of Betty Carter\",\"artist\": \"Betty Carter\",\"price\": 49.99}]'
```
