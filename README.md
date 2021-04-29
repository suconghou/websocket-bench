## websocket 测试工具


bench压测

```
sysctl -w net.ipv4.tcp_fin_timeout=30
sysctl -w net.ipv4.tcp_timestamps=1
sysctl -w net.ipv4.tcp_tw_recycle=1

```

建立一万链接 

./bench -addr ws://192.168.0.5:6060/ws/live1/12 -num 10000

批量测压

```bash
for (( c=0; c<3; c++ ))
do
    docker run --name ws$c -v $(pwd)/bench:/bench -d alpine /bench -addr=ws://192.168.0.5:6060/ws/live$c/1$c -num=30000
done

```

清理

```bash

for (( c=0; c<3; c++ ))
do
    docker kill ws$c
done


for (( c=0; c<3; c++ ))
do
    docker rm ws$c
done

```

static build for linux 

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bench -a -ldflags "-s -w" bench.go
```