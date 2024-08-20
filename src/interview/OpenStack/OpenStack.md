# OpenStack









## 云计算的前世今生
* IT系统架构的演进，分为三个阶段：物理机->虚拟机->云计算
  * 面向物理设备的裸机：应用直接部署运行在物理机，资源利用率很低；
  * 面向资源的虚拟化：虚拟机解决了物理机资源使用率，但是带来了新的问题：管理多台虚拟机；
  * 面向服务的云计算：对虚拟机统一和高效的管理、调度
* 云计算：一种按使用量付费的模式，这种模式提供可用的、便捷的、按需的网络访问，通过互联网进入配置的资源共享池(包括网络、计算、存储、应用软件、服务等)
  * 云计算的几个层次服务
    * SaaS(Software as a Service):软件即服务，把在线软件作为一种服务(直接可使用的数据库服务)；
    * PaaS(Platform as a Service):平台即服务，把平台作为一种服务； 
    * IaaS(Infrastructure as a Service)：基础设施即服务，把硬件设备作为一种服务(openstack)；
* OpenStack：由Rackspace和NASA共同开发的云计算平台，是一个开源的IaaS云计算平台(因为AmazonCloud非常火，且付费)，任何人可以自行建立和提供云端运算服务，每半年发布一次，用Python编写。
  * 国内公有云：阿里、腾讯、金山、华为云；国内基于OpenStack的私有云：浪潮、九州、易捷行云(EasyStack)、海云捷迅、青云；基于docker、k8s二次研发的有灵雀云、时速云；
* 社区与链接
  * [社区](https://www.openstack.org)  wiki.openstack.org
  * 邮件列表：
    * http://wiki.openstack.org/MailingLists#General_List   
    * http://wiki.openstack.org/MailingLists#Development_List
    * http://wiki.openstack.org/MailingLists#Operators
  * [如何贡献代码](http://wiki.openstack.org/HowToContribute)
  * [源代码管理](http://wiki.openstack.org/GerritWorkflow)
  * [文档](https://docs.openstack.org/2024.1/)
## OpenStack
OpenStack为私有云和公有云提供可扩展的弹性云计算服务，这种服务云必须时简单部署并且扩展性强。
1. 模块松耦合
2. 组件配置灵活
3. 二次开发容易

* 基础架构
  ![img.png](img.png)

* 管理界面
  ![img_1.png](img_1.png)

* 构成组件
  ![img_2.png](img_2.png)

* OpenStack共享服务组件
  * 数据库服务(DB Service): **MariaDB** & MongoDB
  * 消息传输(Message Queue)：**RabbitMQ**
  * 缓存(Cache)：Memcached、时间同步：**NTP**
  * 存储(Storage Provider):**ceph**、GFS、LVM、ISICI等
  * 高可用即负载均衡：pacemaker、HAproxy、**keepalive**、**IVS**等
* OpenStack核心组件
  * 身份服务(Identity Service)：**Keystone**
  * 计算(Computer)：**Nova**
  * 镜像服务(Image Service)：**Glance**
  * 网络&地址管理(Network)：**Neutron**
  * 对象存储(Object Storage)：Swift
  * 块存储(Block Storage)：Cinder
  * UI界面(Dashboard)：Horizon
  * 测量(Metering)：**Ceilometer**
  * 部署编排(Orchestration)：Heat

## OpenStack常用组件及服务

### [共享组件-NTP](https://www.cnblogs.com/cloudhere/p/10673458.html)

* 标准时区：
  * 地球分为东西十二个区域，共计 24 个时区
  * 格林威治作为全球标准时间即 (GMT 时间 )，东时区以格林威治时区进行加，而西时区则为减。
  * 地球的轨道并非正圆，在加上自转速度逐年递减，因此时间就会有误差在计算时间的时，
    最准确是使用“原子震荡周期”所计算的物理时钟。这种时钟被称为标准时间即— Coordinated Universal Time(UTC)
  * 随着时间的误差，有些工作是无需进行时间精确即可以完成。但有些工作就必须精确时间从而可以完成目标任务。
  * 因此时间的同步有了需求。目前所使用的就是 Network Time Protocol 协议。即网络时间协议。
  ```shell
  date +"%F %T"
  ```
  
* NTP同步时钟服务
  * NTP工作请求
    * 客户端采用随机端口向NTP服务器(UDP:123)发出时间同步请求
    * NTP服务器收到请求后发出调校时间
    * NTP客户端收到NTP服务器的消息后，进行调整，完成时间同步
    ```shell
      自己服务器 -> 上一级的时间服务器 -> 上上一级时间服务器 -> ... -> 根时间服务器
    ```
  * 同步服务器时间方式有两种
    * 一次性同步(手动同步)：`ntpdate + 时间服务器的域名或ip地址`,[ip地址查看](http://www.ntp.org.cn/pool)
    ```shell
      ntpdate 120.25.108.11 # 选择阿里云的时间服务器
    ```
    * 服务器自动同步
      * NTP服务器安装 `yum install ntp -y`
      * 查看配置是否存在 `ls -l /etc/ntp.conf`
      * NTP涉及程序 `ntpd、ntpdate、tzdata -update`
      * 相关事件程序 `date、hwclock`
      * 涉及文件 
      ```shell
        /etc/ntp.conf # ntp服务器配置
        /usr/share/zoneinfo/  # 由tzdata所提供的各个时区对应文档
        /etc/sysconfig/clock # 设定时区与是否使用UTC时间
        /etc/localtime # 本地时间文件
      ``` 
      * NTP服务：NTP服务属于C/S架构模式，建立在本地服务时最好与上层服务器进行时间同步来给本地提供时间同步服务
      ```shell
        # 中国NTP服务
        cn.pool.ntp.org
        0.cn.pool.ntp.org
        1.cn.pool.ntp.org
        2.cn.pool.ntp.org
        3.cn.pool.ntp.org
      ```
      * Linux客户端同步
        * 手动同步  `ntpdata 192.168.1.100`
        * 配置文件 ``
        ```shell
          # 编辑配置
          vim /etc/ntp.conf  
          # 配置写入文件
          server 192.168.1.100
          # 启动进程
          system start ntpd
        ```
      * 查看上层NTP服务状态 `ntpq -p`

### [共享组件-RabbitMQ](https://www.cnblogs.com/cloudhere/p/10673654.html)

* MQ(Message Queue):消息队列,是一种应用程序间通信的方法，消息传递是通过消息中间件来交互实现，而非直接调用来通信，有效降低了系统的耦合性。
* AMQP(Advanced Message Queuing Protocol)：高级消息队列协议，是应用层协议的一个开放标准，为面向消息的中间件设计。AMQP的主要特征是面向消息、队列、路由(包括点对点和发布/订阅)、可靠性和安全。
* RabbitMQ：属于一个流行的开源消息队列系统。属于AMQP(高级消息队列协议)标准的一个实现。是应用层协议的一个开放标准。可用于分布式系统中存储转发消息，易用、可扩展、高可用。
  * Erlang编写
  * 支持持久化
  * 支持HA
  * 提供C#、erlang、Java、Perl、Python、Ruby、Golang等client的客户端
* 耦合与解耦
  * 耦合
    * 指两个或两个以上的体系或两种运动形式间相互作用而彼此影响以致联合起来的现象；
    * 软件工程中，对象的耦合度就是对象之间的依赖性，对象耦合程度越高，维护成本也越高，所以设计时组件之间尽量降低耦合性；
  * 解耦
    * 即解除耦合关系
    * 软件工程中，降低系统耦合度即可理解为解耦，模块之间有依赖就必然存在耦合；理论上的零耦合做不到，只能通过设计降低系统间的耦合关系；
    * 核心设计思想：尽可能减少代码耦合。让数据模型、业务逻辑和视图显示三层之间彼此降低耦合，把关联依赖将至最低。
* RabbitMQ中名词解释
  * broker：消息队列的服务器实体
  * exchange：消息交换机，消息按照指定的规则，路由到对应的队列；
  * queue：消息队列载体，存放消息(实际在segment中)，每个消息会被投递到一个或多个队列中。
    * Binding：绑定，作用是把exchange和queue按照路由规则绑定起来；
    * RoutingKey：路由关键字，exchange根据这个关键字进行消息投递(还和exchange的类型有关)
  * vhost：虚拟主机，一个broker可以设置多个vhost，用作不同用户的权限分离
  * producer：生产者，投递消息的程序；
  * consumer：消费者，消费消息的程序；
  * channel：消息通道，在客户端的每个连接里，可以建立多个channel，每个channel代表一个会话任务；
* RabbitMQ工作原理
MQ是消费-生产者的典型实现，一端不断往消息队列中写入消息，另一端则通过读取或者订阅队列中的消息。这种异步处理降低了系统的响应时间，增大了系统的吞吐量。
  * 客户端(生产者/消费者)连接到broker，打开一个channel；
  * 客户端(生产者/消费者)申明一个exchange(Fanout、Direct、Topic、Header)，并设置相关属性
  * 客户端(生产者/消费者)申明一个queue，并设置相关属性
  * 客户端(生产者/消费者)使用routingKey，在exchange和queue间建立好绑定关系；
  * 客户端(生产者/消费者)投递/消费消息
  * 生产者：exchange接收到消息后，就根据消息的key和已设置的binding关系，进行消息路由，投递到对应的一个或多个队列中
  * 消费者：从exchange中pull消息；订阅的话等待消息从MQ中push过来；
* RabbitMQ的metadata
  * 元数据可持久化在RAM(内存)或Disk(磁盘)上，从这个维度来讲，RabbitMQ集群的节点分为RAM Node和Disk Node
    * RAM Node：元数据存放在RAM中；
    * Disk Node：元数据持久化到磁盘
  * 单节点只允许Disk Node，否则一旦重启所有数据都会丢失，可在集群环境中选择哪些是RAM Node。
* RabbitMQ的集群部署。。。

### [共享组件-MemCached缓存](https://www.cnblogs.com/cloudhere/p/10673794.html)

* 缓存系统
  * 静态页面：主要使用Nginx或者CDN做静态页面缓存
  * 动态页面：前端使用Nginx或者CDN做缓存，后端大多数使用关系型数据库，关系型数据库的并发比较低，可以提前将数据写入缓存系统(Redis、Memcache等)中，提高系统的并发能力。
  * 通过在内存中缓存数据和对象来减少读取数据库的次数，从而提高网站的访问速度，加速动态Web应用、减轻数据库负载。
* Memcached概念
  * Memcached是一个开源的、高性能的分布式内存对象缓存系统，竞争对手是Redis这种。
  * Memcached把经常需要存取的对象或数据缓存在内存中，数据通过API访问，再经过HASH之后存放到位于内存上的hash表内，以key-value的方式存放。Memcached没有实现访问认证及安全管理控制，所以需要部署在安全的位置。
  * Memcached节点的物理内存不足时，会使用LRU算法来淘汰不活跃的数据。
  * Memcached易于二次开发，所以也有一定适用范围。此外还提供了很多语言开箱即用的客户端。
  * OpenStack的KeyStone身份认证中，使用Memcached来缓存租户的Token等身份信息；在Horizon和对象存储Swift项目也有利用到Memcached来缓存数据。
* Memcached缓存流程
  * 检查客户端请求的数据是否在memcached缓存中，如果存在则直接将数据返回；
  * 如果不存在，则去数据库中查询，查询到的数据返回给客户端，同时把数据缓存一份到memcached中；
  * 每次更新数据库的同时更新memcached中的数据，确保数据的一致性；
  * 当memcached的内存用完后，会根据LRU算法来淘汰失效的数据；
* Memcached功能特点
  * 协议简单：基于文本行协议，能直接通过Telnet在Memcached服务器上存取数据
  * 基于libevent的异步事件处理：libevent是利用C语言开发的程序库，将BSD系统的kqueue、Linux的Epoll等事件处理封装成一个接口，确保及时服务器端的连接数，Memcached利用此库进行异步事件处理。
  * 内置的内存管理方式：Memcached有一套自己的内存管理方式(Redis使用Meta和Google的)，这套方式非常高效，所有数据都保存在Memcached的内存中；当内存占满时，使用LRU算法淘汰过期数据。
    但是Memcached不考虑容灾，不会进行持久化，一旦断电或者重启，所有数据都会丢失。区别于Redis的持久化。
  * 节点互相独立的分布式：各个Memcached服务器之间互不通信，都独立存储数据，不共享任何信息。通过对客户端的设计，让Memcached具有分布式、能支持海量缓存和大规模应用。
* 使用Memcached应该考虑的因素
  * Memcached服务单点故障：集群每个节点独立存储数据，不同步数据，所以如果一个Memcached节点故障或重启，该节点缓存的数据全部丢失，再次访问时需要再次缓存到该服务器；
  * 存储空间限制：数据存储到内存中，会受到寻址空间大小限制，32位系统缓存大小为2G，64位系统理论是无限的，不过要看分配给Memcached服务器的物理大小；
  * 储存单元限制：数据以key-value形式存储，能存储的数据key大小为250字节，value大小为1MB，超过这个值不允许存储；
  * 数据碎片：内存存储单元室按照Chunk来分配的，但是所存储的value数据大小不一定正好等于一个Chunk的大小，所以必然会造成内存碎片，浪费存储空间；
  * LRU算法局限性：LRU算法并不是针对全局空间的存储数据，而是针对Slab，Slab是Memcached中具有同样大小的多个Chunk集合；
  * 数据访问安全性：Memcached服务端并没有相应的安全认证机制，通过非加密的telnet连接即可对Memcached服务器端的数据进行各种操作。








## 参考链接
* [云计算OpenStack](https://www.cnblogs.com/cloudhere/category/1439151.html)









