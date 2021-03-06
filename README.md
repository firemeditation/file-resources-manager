# 图书数字资源管理 #

--------------

**特别注意：**此项目处于开发阶段。

**所需组件：**

本人特别喜欢使用较新的东西，所以目前来看是这样的：

+ 开发语言为Golang，在Go 1.2下编译通过；

+ 数据库采用PostgreSQL 9.3+，并使用了github.com/lib/pq这个数据库驱动；

+ 使用了github.com/msbranco/goconfig作为配置文件的读取。

大概就是这样吧，如果少了哪个包，在编译的时候会有提示，所以也就是无所谓了。

-----

## 主要特性 ##

+ Golang语言开发，具有良好的跨平台性，初步测试，操作系统支持Linux与Windows，处理器架构支持x86与ARM。

+ Server/Client/Web混合模式。

	+ 用户所有操作均通过Web界面完成，一个界面搞定一切，不需要切换操作界面。

	+ 完全无刷新的Web界面，用户感觉如同在操作本地应用程序。

	+ 对于文件的上传下载，交由Client处理，避免HTTP协议传输文件的诸多瓶颈，用户使用时处于后台状态，用户通过Web界面操作来控制Client调用。

+ 文件传输支持多线程以及资源读写锁，保证传输效率和安全性。

+ Server端文件保存支持多盘多区，且自动分布存放，减少写入延时。

+ 支持多机构管理。

+ 异步缓存式搜索（异步缓存全文索引），减少搜索时的系统开销。 _[基本实现（不涉及碎片化编辑的部分）]_

+ 碎片化编辑，可以在线新建、编辑图书各章节，并嵌入媒体文件。 _[暂不实现]_

----

## 编译与使用 ##

在PostgreSQL中建立一个新的数据库，并将docs/database.sql文件导入进数据库。

通过buildFRM脚本文件，可以一次性编译客户端和服务器端，生成的文件为frmClient和frmServer，连同两个配置文件以及Web界面所需静态文件static/路径，均保存在bin/路径下。

如果跨平台编译，请查看buildFRM文件中的被注释语句。

按需要修改两个配置文件里的相关参数，特别是和数据库相关的以及几个端口或地址配置。

在服务器上启动frmServer，在个人计算机上启动frmClient。

在个人计算机上通过浏览器（目前只测试了Google Chrome）访问服务器端地址，例如：http://mydomain.com:9998 。

使用默认用户名密码登录，用户名：root，密码：123456。

在Web界面右上角检查是否与客户端正常连接。

如果一切正常则可以开始使用。

------

## 演示视频 ##

“图书数字资源管理”最新功能演示（2014年2月21日）  [vimeo](https://vimeo.com/87276191)  [viddler](http://www.viddler.com/v/af309ebc)

