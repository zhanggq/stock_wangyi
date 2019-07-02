## 简介

该工程主要作用是调用网易财经接口获取股票日线数据

- 后端使用Go语言开发，基于Beego框架（主要用了Route和DB）

- 前端使用Amaze框架

- 数据来源使用网易财经数据

- 数据库使用mysql

- 打包使用go镜像来封装

- 工程运行后就会自动开始从网易财经爬取股票从2010年至今的日线数据

- 默认每日晚上10点开始更新当日数据

## 前序准备

- 需要一台Linux服务器，配置任意

- 服务器上需要安装Docker

- 服务器上需要安装Mysql。或者仅安装Mysql-client，Server使用容器来启动（可以参考我的[mysql5.7](https://github.com/zhanggq/mysql5.7)）

## 镜像制作

在服务器上执行下列命令：

```Bash
# git clone https://github.com/zhanggq/stock_wangyi.git
# cd stock_wangyi
# sh build.sh
```

注：考虑到国内经常go get失败，所以依赖包都已经放在vendor目录下了，所以工程有点点大

## 配置说明

配置文件位于PPGo_amaze\conf\app.conf，主要涉及web启动的端口号和Mysql配置


## 启动Mysql

假设使用我的[mysql5.7](https://github.com/zhanggq/mysql5.7)，可以按照如下方式启动

在服务器上执行下列命令：

```Bash
# mkdir -p /home/lib/mysql/wangyi
# docker run -d -p 13306:3306 --name mysqlv5.7 -P -e mysqld -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=testDb -v /home/lib/mysql/wangyi:/var/lib/mysql mysql:v5.7
```

## 准备库与用户

在服务器上执行下列命令：

```Bash
# mysql -uroot -p123456 -h 127.0.0.1 -P13306 -e "create database hcs character set utf8 collate utf8_bin;"
# mysql -uroot -p123456 -h 127.0.0.1 -P13306 -e "create database tushare character set utf8 collate utf8_bin;"
# mysql -uroot -p123456 -h 127.0.0.1 -P13306 -e "grant all privileges on *.* to hcs@localhost identified by 'hcs' with grant option;"
# mysql -uroot -p123456 -h 127.0.0.1 -P13306 -e "grant all privileges on *.* to 'hcs'@'%' identified by 'hcs' with grant option;"
# mysql -uroot -p123456 -h 127.0.0.1 -P13306 -e "flush privileges;"
```

## 导入数据

在服务器上执行下列命令：

```Bash
# cd stock_wangyi/sql
# for SQL in *.sql; do mysql -uroot -p123456 -h 127.0.0.1 -P13306 hcs < $SQL; done
```

## 检查导入结果

在服务器上执行命令：

```Bash
# mysql -uroot -p123456 -h 127.0.0.1 -P13306 hcs  -e "show tables"
```

## Mysql容器验证通过后，启动stock容器

在服务器上执行命令：

```Bash
# docker run --privileged -tid -h wangyi --name=wangyi --net=host --restart=always stock_wangyi:latest
```

## 进入容器检查日志

在服务器上执行进入容器命令：

```Bash
# docker exec -it XXXX bash
```

进入容器后执行：

```Bash
# tail -f /PPGo_amaze/info.log
```

## 登入Web页面

登入自己的服务器Ip+8010端口http://xxx.xxx.xxx.xxx:8010/。用户为admin 密码为admin@123

![](https://i.imgur.com/BcNzBAS.jpg)

登入后进入首页，这个页面抄来后没改过，没有实际意义，有数据的为图中红色框“深证”，“创业”，“个股”

![](https://i.imgur.com/T4caZzs.png)

点击“深证”后展示深证k线

![](https://i.imgur.com/XQemBpw.png)

## 数据库说明

“深证”，“创业”，“上证”的股票代码分别保存在pp_name_000，pp_name_300，pp_name_600表中。所有的股票日线数据保存在pp_value表中。

![](https://i.imgur.com/6ukKxJp.png)

注1：由于本人不爱炒上证股票，固代码中没有爬取上证股票日线数据；
注2：作者自己腾讯云的服务器，第一次跑需要4小时才能同步完所有的数据。
