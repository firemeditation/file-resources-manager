# 图书数字资源管理 #

--------------

**特别注意：**此项目处于开发阶段，功能完成度非常低，暂无使用价值。

**所需组件：**

本人特别喜欢使用较新到东西，所以目前来看是这样的：

+ 开发语言为Golang，在Go 1.2下编译通过；

+ 数据库采用PostgreSQL 9.3+，并使用了github.com/lib/pq这个数据库驱动；

+ 使用了github.com/msbranco/goconfig作为配置文件的读取。

大概就是这样吧，如果少了哪个包，在编译的时候会有提示，所以也就是无所谓了。

----

## 编译与使用 ##

在PostgreSQL中建立一个新的数据库，并将docs/database.sql文件导入进数据库。

通过buildFRM脚本文件，可以一次性编译客户端和服务器端，生成的文件为frmClient和frmServer，连同两个配置文件以及Web界面所需静态文件static/路径，均保存在bin/路径下。

如果跨平台编译，请查看buildFRM文件中的被注释语句。

按需要修改两个配置文件里的相关参数，特别是和数据库相关的以及几个端口或地址配置。

在服务器上启动frmServer，在个人计算机上启动frmClient。

在个人计算机上通过浏览器（推荐Google Chrome）访问服务器端地址，例如：http://mydomain.com:9998。

使用默认用户名密码登录，用户名：root，密码：123456。

在Web界面右上角检查是否与客户端正常连接。

如果一切正常则可以开始使用。
