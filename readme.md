# Golang 实现的通过 QQWry.dat 查找 IP 归属地的库

# 安装
`go get -u github.com/elsejj/qqwry`

# 使用

## 使用 db

```Go
package main
import "github.com/elsejj/qqwry"

func main() {
	path := "qqwry.dat"
	ip   := "183.224.52.133"
	db, err := qqwry.NewDb(path)
	if err != nil {
		country, area := db.Search(ip)
		//TODO: use country area
	}
}
```

## 使用 http 服务

```Go
package main
import "github.com/elsejj/qqwry"

func main() {
	path := "qqwry.dat"
	addr := "127.0.0.1:8000"
	qqwry.StartHttp(addr, path)
}
```

则通过 [http://127.0.0.1:8000/?ip=183.224.52.133&ip=xxx.xx.xxx.xx](http://127.0.0.1:8000/?ip=183.224.52.133&ip=202.123.44.23)

# 性能测试

## Go bencn

`BenchmarkSearch  1000000              1028 ns/op`

## ab

```
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /?ip=183.224.52.133&ip=10.12.23.44
Document Length:        159 bytes

Concurrency Level:      800
Time taken for tests:   4.973 seconds
Complete requests:      80000
Failed requests:        0
Write errors:           0
Keep-Alive requests:    80000
Total transferred:      24080903 bytes
HTML transferred:       12720477 bytes
Requests per second:    16085.78 [#/sec] (mean)
Time per request:       49.733 [ms] (mean)
Time per request:       0.062 [ms] (mean, across all concurrent requests)
Transfer rate:          4728.52 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     0   49   8.7     47     144
Waiting:        0   49   8.7     47     144
Total:          0   49   8.7     47     144

Percentage of the requests served within a certain time (ms)
  50%     47
  66%     50
  75%     52
  80%     53
  90%     59
  95%     63
  98%     67
  99%     71
 100%    144 (longest request)
 ```
