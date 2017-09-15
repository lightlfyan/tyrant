# tyrant 分布式压测


###  如何部署
1. 启动master
```
nohup ./master &
```
2. 启动cell
```
nohup ./cell &
```
3.其他说明
```
master只有一台, http监听端口8001, cell长连接端口9999
cell可多台, 每个ip只能启动一个cell进程, cell启动可通过 -addr 来指定master地址(-addr "127.0.0.1:9999")
```

### 架构设计

![arch](http://127.0.0.1/arch.png)

>协议

    protocolbuf


>网络拓扑

    server-client 模式


>任务流程

    cell启动后主动向master进行注册(自动重连)
    master分配任务到cell, 当cell的任务完成后, 会向master发送压测结果,
    master收到所有cell返回的任务报告后, 进行一次汇总, 把数据存入内存, 以及本地的文件数据库
    
    
 # 使用说明
 
 
 ## 任务创建页面 Create
 

 >url
 
 1. 被压测的url, 例如 *https://www.example.com*
 2. url可以使用逗号分割, 这样压测工具会随机选取url来进行访问, 例如: http://www.example.com;https://10.102.10.22/api
 3. url支持可变参数, 需要在upload页面上传文件, 各式参考下面:
 
 ```csv
 abcd,1234,5678
 sowl,2123,1425
 
 每条数据使用逗号分隔,分别对应标记 {{.P1}} {{.P2}} {{.P3}}
  
 例如
 http://127.0.0.1?openid={{.P1}}&username={{.P2}}&passwd={{.P3}}
```


>num 

请求的总数量, 如果有1000个请求, 压测工具把请求平均分配给压测机

>conc

并发数量

> qps

暂不启用

>gzip

勾选,则开启gzip请求,服务器返回数据会进行gzip压缩

>method

get post

>timeout

超时时间

>accept

http 请求接受的格式 Accept-Content

> contenttype

http body的格式 Content-Type

> keepAlive

勾选,则http请求复用tcp连接, qps会变高,但不符合真实场景

>host

http 请求的域名, 可以url里填写ip, 然后host填写域名

>cookie


http cookie

```
例如:
username=abc; passwd=123;
```


>headers

http请求头

```bash
例如:
Accept-Encoding=gzip&Content-Type=text/plain
```
