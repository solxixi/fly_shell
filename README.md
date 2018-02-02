# fly_shell

交互式并发执行远程命令的运维小工具，可根据需求轻易扩展
###应用场景

- 远程多台主机执行命令
- 远程一次查询多台主机log并合并输出，可以配合websocket进行主机或docker日志的实时查询

###安装依赖

go get golang.org/x/crypto/ssh

###用法

./fly_shell -h "10.0.0.1,10.0.0.2" -c "tail -f /data/logs/xxx.log"
