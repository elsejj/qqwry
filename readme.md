# Golang 实现的通过 QQWry.dat 查找 IP 归属地的库

# 安装

## 命令行使用

```bash
go install github.com/elsejj/qqwry@latest
```

## 作为库使用

```bash
go get -u github.com/elsejj/qqwry/qqwry
```

请查看 [cmd](/cmd) 目录下的代码示例。

# 命令行功能


## 查询 IP 地址

### 直接查询 IP 地址的归属地信息。
```bash
qqwry search 183.224.52.133 115.60.135.120
```

### 查询 IP 地址的归属地信息，并输出为 JSON 格式。
```bash
qqwry search -f json 183.224.52.133 115.60.135.120
```

### 将给定文件中的 IP 替换为归属地信息。

```bash
qqwry search datas/visit.txt
```

### 将给定文件中的 IP 替换为归属地信息 (输出为 JSON 格式)。
```bash
qqwry search -f json datas/ips.json
```

## 更新数据库

```bash
qqwry update
```

这将从 [qqwry.dat](https://github.com/metowolf/qqwry.dat) 下载最新的 IP 数据库

感谢 **metowolf** 的更新


## 作为服务

可以使用 `qqwry serve` 命令启动一个 HTTP 服务，提供 IP 查询功能。


### 以 GET 请求查询 IP 地址的归属地信息。

```bash
curl "http://localhost:11223/?ip=183.224.52.133&ip=192.168.1.2"
```

### 以 POST 请求查询 IP 地址的归属地信息。

```bash

curl http://127.0.0.1:11223/ -H 'content-type: application/json'  -d '{"userIP":"183.224.52.133","localIP":"192.168.1.2"}'
```

