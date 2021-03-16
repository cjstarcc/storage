## 分布式对象存储的管理方式

对象存储是以对象来管理数据的，一个对象包含了三个部分：对象的数据，对象的元数据，和对象的唯一标识符（对象ID）

对象数据：就是对象存储的数据本身，比如xxx.avi,xxx.jpg等等

对象的元数据：就是对象存储的描述信息，比如xxx.avi对象的文件名，大小，创建时间，内容，演员是哪个老师等等

对象唯一标识符：用于引用对象，具有全局唯一性，可以使MD5或者是其他的散列值等等



## 访问数据的方式

此程序使用REST风格的HTTP协议来访问。



## 对象存储的优势

由于传统的存储是存储在单机上的，会造成数据越多开销会呈指数级的上升。而对象存储的会只需要增加节点就行。


## 安装
### 1.创建存储目录
```shell script
mkdir /data
```
### 2.安装rabbitmq
```shell script
sudo apt-get install rabbitmq-server # 安装rabbitmq服务器
sudo rabbitmq-plugins enable rabbitmq_management # 安装rabbitmq客户端管理工具
wget 127.0.0.1:15672/cli/rabbitmqadmin # 下载管理工具
python3 rabbitmqadmin declare exchange name=apiServers type=fanout # 创建apiServer exchange
python3 rabbitmqadmin declare exchange name=dataServer type=fanout # 创建dataServer exchange
sudo rabbitmqctl add_user test test # 添加test用户
sudo rabbitmqctl set_permissions -p / test ".*" ".*" ".*" #为test用户增加所有权限
```

### 3. 安装elasticsearch 7.11
```shell script
wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.11.2-linux-x86_64.tar.gz
tar -xzvf elasticsearch-7.11.2-linux-x86_64.tar.gz
adduser es #添加es用户
chown -R es:es elasticsearch-7.11.2
su es
/elasticsearch-7.11.2/bin/elasticsearch -d
```
