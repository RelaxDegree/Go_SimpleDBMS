## 分为Server端和Client端的DBMS demo
Client 端简介 ： 具有通过cli命令行与Server端交互，具有登陆功能和发送SQL命令的功能
client 端
- 首先进行用户权限操作：可以进行登陆/注册操作 同时登陆注册的行为记录到日志中
- Client端需要另起一个文件运行
- 最后进行退出操作或者强制退出时，服务器应该记录这个行为
Server端可以将Client端的命令进行执行，并将所有Client端的命令都记录于日志中，可以进行日志查看
- 当用户端进行登陆或者注册操作时，查询数据库并记录日志。
- 注册成功或者登陆成功时，服务器端记录日志并载入在线用户池

