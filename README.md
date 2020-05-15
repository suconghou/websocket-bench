## websocket 测试工具


bench压测


建立一万链接 

./bench -addr ws://172.168.1.5:6060/ws/live1/12 -num 10000

批量测压

```bash
for (( c=0; c<5; c++ ))
do
    docker run -l ws -v $(pwd)/bench:/bench -d alpine /bench -addr=ws://172.168.1.5:6060/ws/live$c/$c_1 -num=30000
done

```

