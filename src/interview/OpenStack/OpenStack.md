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
  * 国内公有云：阿里、腾讯、金山、华为云等；国内基于OpenStack的私有云：浪潮、九州、易捷行云(EasyStack)、海云捷迅、青云；基于docker、k8s二次研发的有灵雀云、时速云；
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
  ![img_7.png](img_7.png)
* 核心组件交互逻辑
  ![img_8.png](img_8.png)

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

### [核心组件-Keystone身份认证服务](https://www.cnblogs.com/cloudhere/p/10811666.html)
  ![img_6.png](img_6.png)
  Keystone是opensta的组件之一，是用于为OpenStack中的其他组件提供统一认证的服务，包括身份验证、令牌发放和校验、服务列表、用户权限定义等等。
  OpenStack云环境中的所有服务之间的授权和认证都需要经过keystone，所以keystone是云平台中第一个需要安装的服务。
* Keystone简介
  * Keystone作为OpenStack的基础服务，功能职责有以下：
    * 管理用户及其权限
    * 维护OpenStack Service的Endpoint
    * Authentication(认证)和Authorization(授权)
  * 理解以下概念(User、Role、Credentials、Authentication、Token、Project、Service、Endpoint)
    * User(用户)
      * 指代任何使用OpenStack的实体，可以是真正的用户，也可以是其他系统或者服务；
      * 当user请求访问OpenStack时，Keystone会对其进行身份验证。管理在：Horizon-> Identity -> Users管理界面；
        * admin：超级管理员，负责OpenStack服务的管理和访问权限
        * demo：常规(非管理)用户，任务应该使用无特权的项目和用户，所以需要创建demo项目和demo用户。
        * 除admin和demo外，OpenStack也为nova、cinder、glance、netron等服务创建了对应的User，admin可以管理这些User。
    * Credentials(证书、凭证)：是User来证明自己身份的信息，可以是
      * 用户名+密码；
      * Token
      * API Key
      * 其他高级方式
    * Authentication(验证、认证)：是Keystone验证user身份的过程。User访问OpenStack时向Keystone提交用户名和密码形式的Credentials，Keystone验证通过后会给user签发一个Token座位后续访问的credentials。
    * Token(令牌)：由数字和字母组成的字符串，User成功Authentication后Keystone生成Token并返回给user
      * Token用作访问Service的credentials凭证；
      * Service会通过Keystone验证Token的有效性；
      * Token的有效期默认24H；
    * Project(项目)：用于将OpenStack的资源(计算、存储和网络等)进行分组和隔离；根据OpenStack服务对象不同，Project也可以是一个客户(公有云，也叫租户)、部门或项目组(私有云)。注意：
      * 资源的所有权是属于Project的，而不是User；
      * 在OpenStack的界面和文档中，Tenant/Project/Account这几个术语是通用的，但是倾向于使用Project；
      * 每个user(包括admin)必须挂在Project里才能访问Project的资源；User和Project是多对多关系；
      * admin==root，拥有最高权限；
      * Horizon-> Identity -> Project中管理Project；
    * Service(服务)：包括Computer(Nova)、BlockStorage(Cinder)、ObjectStorage(Swift)、Image Service(Glance)、Network Service(Neutron)等；
      每个Service提供若干个Endpoint，User通过Endpoint对资源进行访问和操作；
    * Endpoint(端点)：是网络中一个可访问的地址，通常是一个URL。Service通过Endpoint暴露自己的Api。Keystone负责管理和维护每个Service的Endpoint。
    * Role(角色)：安全包含两部分：认证(Authentication)和鉴权(Authorization)
      * 认证(Authentication)：解决`who you are`
      * 鉴权(Authorization)：解决`what can you do`
      * keystone借助Role实现Authorization:
        * keystone创建Role角色；
        * 为每个user绑定一个或多个角色；
        * 用户通过Endpoint来访问或操作资源时，Keystone根据预先分配的角色，判断本次操作是否有对应的权限。
* Keystone基础架构
![img_3.png](img_3.png)
  * Token：用来生成和管理Token
  * catalog：用来储存和管理service/endpoint；
  * identity：用来管理tenant/user/role和验证
  * policy：用来管理访问权限
* 配置文件：/etc/keystone/keystone.conf
  ```shell
    [DEFAULT]
    [assignment]
    [auth]
    [cache]
    [catalog]
    [cors]
    [cors.subdomain]
    [credential]
    [database]
    connection = mysql+pymysql://keystone:KEYSTONE_DBPASS@controller/keystone
    [domain_config]
    [endpoint_filter]
    [endpoint_policy]
    [eventlet_server]
    [federation]
    [fernet_tokens]
    [healthcheck]
    [identity]
    [identity_mapping]
    [kvs]
    [ldap]
    [matchmaker_redis]
    [memcache]
    [oauth1]
    [oslo_messaging_amqp]
    [oslo_messaging_kafka]
    [oslo_messaging_notifications]
    [oslo_messaging_rabbit]
    [oslo_messaging_zmq]
    [oslo_middleware]
    [oslo_policy]
    [paste_deploy]
    [policy]
    [profiler]
    [resource]
    [revoke]
    [role]
    [saml]
    [security_compliance]
    [shadow_users]
    [signing]
    [token]
    provider = fernet
    [tokenless_auth]
    [trust]
  ```

### [核心组件-Glance镜像服务](https://www.cnblogs.com/cloudhere/p/10811888.html)
* Glance介绍：Glance是OpenStack项目中负责镜像管理的模块，功能包括虚拟机镜像的查找、注册和检索等。提供restfulAPI可以查询虚拟机镜像的metadata以及获取惊喜。
  Glance可以将镜像保存到多种后端存储上，如简单的文件存储或对象存储。
  * Image(镜像)：是可执行文件的系统快照。
  * Image Service(镜像服务)：功能是管理Image，让用户能够对镜像做增删改查。OpenStack使用Glance；公共比较普遍是docker；
    * 提供REST API让用户查找获取image的metadata和image本身；
    * 支持多种方式存储，包括普通文件系统、Swift、Amazon S3等；
    * 对instance执行snapshot创建新的image；
* Glance架构：
  * Glance-Api：系统后台运行的服务进程。**对外提供REST API**，相应image的查询、获取和存储；注意glance-api不会真正处理请求。如果操作metadata，会转发到glance-registry；操作image本身，则会转发到store backend；
  * Glance-Registry：后台服务进程，**负责处理和存储image metadata**；
  * Database：数据库存储metadata等，默认是MySQL；
  * Store Backend：真正存放image的地方。支持多种Backend：
    * A directory on a local file system(默认配置-文件系统)
    * GridFS
    * Ceph RDB
    * Amazon S3
    * Sheepdog
    * OpenStack Block Storage(Cinder-块存储)
    * OpenStack Object Storage(Swift-对象存储)
    * VMware ESX
    * 具体使用哪种，在`/etc/glance/glance-api.conf`中配置
* Glance创建镜像
  * Web UI(Horizon Dashboard)：界面创建
  * CLI：命令行创建(比Web功能更全更高效)
* 配置文件
  * `/etc/glance/glance-api.conf`
  * `/etc/glance/glance-registry.conf`

### [共核心组件-Nova计算服务](https://www.cnblogs.com/cloudhere/p/10811917.html)
* Nova介绍：是OpenStack最核心的服务，负责维护和管理云环境的计算资源。OpenStack作为IaaS的云操作系统，虚拟机生命周期管理也就是通过Nova来实现的。
  * 功能及用途
    * 实例生命周期管理
    * 管理计算资源
    * 网络和认证管理
    * REST风格API
    * 异步的一致性通信
    * Hypervisor透明：支持Xen、XenServer/XCP、KVM、UML、VMware vSphere and Hyper-V
  * 架构及组件
    * nova-api：整个Nova的门户，接收和响应客户的REST API调用。对Nova所有的请求首先会经过nova-api处理。在Keystone中查询nova-api的Endpoint；
    * nova-scheduler：虚拟调度服务，负责决定在哪个计算节点上运行虚拟机。创建instance时，用户会提出资源需求(CPU、内存、磁盘等)，OpenStack将这些需求定义在flavor中，用户只需要指定使用哪个flavor就可以了。
      * Filter Scheduler的默认调度器，调度分为两步：
        * 通过过滤器Filter选择满足条件的计算节点(可以使用多个filter依次过滤);
        * 通过权重计算选择在最优的计算节点上创建instance；目前默认是根据计算节点空闲内存量来计算权重值的，空闲内存越多，权重越大；
      * Nova允许使用第三方的scheduler，配置scheduler_driver即可(提现OpenStack的开放性)；
    * nova-compute：管理虚拟机的核心服务，在计算节点上运行。
      * 通过调用Hypervisor API实现节点上的instance的生命周期管理。OpenStack对instance的操作，最后都由nova-compute完成。
        nova-compute和Hypervisor一起实现对instance生命周期的管理，OpenStack中虚拟默认的保存路径：`/etc/lib/nova/instances`
      * 通过Driver架构支持多种Hypervisor：Hypervisor是计算节点上跑的虚拟化管理程序，是虚拟机管理最底层的程序。
        * 不同的虚拟化技术提供自己的Hypervisor，KVM、Xen、VMware等；
        * nova-computer为这些Hypervisor定义了统一的接口，Hypervisor只需要实现这些接口，就可以以Driver的形式即插即用到OpenStack系统中。
    * nova-conductor：nova-compute需要经常更新数据库，如更新和获取虚拟机的状态。处于安全和伸缩性考虑，Computer不直接访问DB，而是委托nova-conductor来实现；
      * 更高的系统安全性
      * 更好的系统伸缩性
    * nova-console：用户可以通过多种方式访问虚拟机的控制台；
    * noca-novncproxy：基于Web浏览器的VNC访问；
    * nova-spicehtml5proxy：基于HTML5浏览器的SPICE访问
    * nova-xvpnvncproxy：基于Java客户端的VNC访问
    * nova-consoleauth：负责对访问虚拟机控制台请求提供Token认证
    * nova-cert：提供x509证书支持
    * Database：Nova会有一些数据需要存放到数据中，一般使用MySQL。数据库安装到控制节点上。Nova使用命名为`nova`的数据库；
    * MessageQueue：Nova使用RabbitMQ，用作各个子服务通信的数据总线。
* Nova组件如何协同工作
  * Nova的物理部署方案：对于Nova，子服务会部署在两类节点上：计算节点和控制节点。计算节点上安装了Hypervisor，上面运行虚拟机，
    * nova-compute需要放在计算节点上
    * 其他子服务则是放在控制节点上
    * 如果控制节点上也有nova-compute，说明该节点既是控制节点，又是计算节点，可以all-in-one。
  * 虚拟机创建Nova子服务过程
    * 客户(用户或其他程序)向API(nova-api)发出请求：`帮我创建一个虚拟机`
    * API对请求做必要校验之后，向MQ发送消息：`让Scheduler创建一个虚拟机`
    * Scheduler从MQ队列中收到`创建虚拟机`请求,然后执行调度算法，从若干计算节点选出节点A；
    * Scheduler向MQ发送一条消息：`在计算节点A上创建这个虚拟机`
    * 计算节点A的Compute(nova-compute)从MQ收到`创建虚拟机`消息，然后在本节点的Hypervisor上启动虚拟机；
    * 在虚拟机创建的过程中，如果Compute需要查询或更新数据库，会通过MQ向Conductor(nova-conductor)发送消息，conductor负责访问数据库。
* Nova创建虚拟机过程
  ![img_4.png](img_4.png)
  * 界面或者命令行通过RESTful API向Keystone获取认证信息
  * Keystone通过用户请求认证信息，并生成auth-token返回给对应的认证请求；
  * 界面或命令行通过RESTful API向Nova-api发送一个boot instance的请求(携带Token)；
  * nova-api接收到请求后先向Keystone发送认证请求，验证Token的有效性和合法性；
  * keystone验证后，如果有效则返回有效的认证和对应的角色；
  * 通过认证后nova-api和数据库通信；
  * 初始化新建虚拟机的数据库记录
  * nova-api通过rpc.call向nova-scheduler请求是否有创建虚拟机的资源；
  * nova-scheduler进程监听消息队列，获取nova-api的请求；
  * nova-scheduler通过查询nova数据库中计算资源的情况，并通过调度算法计算符合虚拟机创建需要的主机
  * 对于有符合虚拟机创建的主机，nova-scheduler更新数据库中虚拟机对应的物理主机信息
  * nova-scheduler通过rpc.cast向nova-compute发送对应的创建虚拟机请求的消息；
  * nova-compute会从对应的MQ中获取创建虚拟机的请求消息；
  * nova-conductor从MQ中拿到nova-compute请求消息，根据消息查询虚拟机对应的信息，然后从数据库中获取虚拟机对应的信息，最后把虚拟机信息通过MQ发送出去；
  * nova-compute从MQ中拿到虚拟机信息。通过keystone的RESTful API拿到认证Token，并通过HTTP请求glance-api获取创建虚拟机所需要的镜像；
  * glance-api又从keystone验证Token的有效性和权限；通过后nova-compute获得虚拟机镜像信息(URL);
  * nova-compute拿到后，通过HTTP请求cinder-api获取镜像；
  * Cinder再次找keystone认证，通过后将镜像资源返回给nova-compute；
  * 此时nova-compute根据instance的信息调用配置的虚拟化驱动来创建虚拟机。
  ![img_5.png](img_5.png)
* 删除nova-compute节点
  * 控制节点上操作查看计算节点`node1`
  ```shell
    openstack host list # 查询节点
    nova service-list
  ```
  * 先将`node1`上的计算服务设置为down，然后disable
  ```shell
    systemctl stop openstack-nova-compute
    nova service-list
    nova service-disable node1 nova-compute
    nova-service-list
  ```
  * 然后可以登录MySQL查看nova-compute的状态，再删除MySQL中的`node1`信息即可

### [核心组件-Neutron网络服务](https://www.cnblogs.com/cloudhere/p/10812189.html)
* Neutron介绍
  * 传统网络管理很大程度依赖管理员手工配置；云环境下网络非常复杂，尤其是多租户场景中，用户随时都会创建、修改和删除网络；所以手工配置无法满足。
    云时代灵活性和自动化成为主流：软件定义网络SDN(Software-Defined Networking)。
  * Neutron的设计目标是实现“网络即服务-Networking as a Service”.所以在设计上遵循了基于SDN实现网络虚拟化的原则，再实现上充分利用Linux系统的各种网络相关技术。
  * OpenStack网络中的SDN组件叫Quantum，因为版权更名为Neutron。
* 网络基本概念
  * network：是一个隔离的二层广播域，Neutron支持多种类型的Network，包括：local、flat、VLAN、VxLAN和GRE；
  * subnet：IPv4或IPv6地址段，instance的IP从subnet中分配、每个subnet需要定义IP地址的范围和掩码；network和subnet是一对多关系；
  * port：虚拟交换机上的一个端口。port上定义了Mac地址和IP地址，当instance的虚拟网卡VIF绑定到port时，port会将Mac和IP分配给VIF；subnet与port时一对多关系。
  * 小结：下面总结 Project，Network，Subnet，Port 和 VIF 之间关系。
    Project 1 : m Network 1 : m Subnet 1 : m Port 1 : 1 VIF m : 1 Instance
* Neutron功能
  * neutron为整个OpenStack环境提供网络支持，包括二层交换、三层路由、负载均衡、防火墙和VPN等；
    * 二层交换switching
    * 三层路由routing
    * 负载均衡load-balancing
    * 防火墙firewalling
      * Security Group：通过iptables限制进出instance的网络包
      * Firewall-as-a-Service：FWaaS，限制进出虚拟机路由器的网络包，也是通过iptables实现。
  * 优点：OpenStack中的SDN组件架构也属于可插拔类型。通过各种插件可以管控不同种类的交换机、路由器、防火墙、负载均衡并实现FWaaS等许多功能。
* Neutron部署方案
  * 方案一：控制节点+计算节点
    * 控制节点：部署的服务包括：neutron server、core plugin的agent和Service plugin的agent
    * 计算节点：core plugin的agent，负责提供二层网络功能；
    * 注意：
      * core plugin和Service plugin已经集成到neutron server中，不需要独立运行plugin服务；
      * 控制节点和计算节点都需要部署core plugin的agent，因为通过该agent控制节点与计算节点才能建立二层连接；
      * 可以部署多个控制节点和计算节点。
  * 方案二：控制节点+网络节点+计算节点
    * 控制节点：部署neutron server服务
    * 网络节点：部署的服务包括：core plugin的agent和Service plugin的agent
    * 计算节点：部署core plugin的agent，负责提供二层网络功能
    * 注意：
      * 控制节点只负责通过neutron server响应API请求；
      * 网络节点实现数据交换、路由以及Load balance等高级网络服务
      * 可以增加网络节点承担更大的负载；
      * 可以部署多个控制节点、网络节点、计算节点
      * 该方案特别适合规模较大的OpenStack集群
* Neutron多个网卡区分不同类型的网络数据
  * Management网络：用于节点之间MQ内部通信已经访问DB服务，所有节点都需要连接到management网络；
  * API网络：OpenStack各组件通过该网络向用户暴露API服务。Keystone、Nova、Neutron、Glance、Cinder、Horizon的endpoint均配置在API网络上；
  * VM网络：又叫tenant网络，用于instance之间通信；类型包括(local、flat、vlan、vxlan和gre)；VM网络由neutron配置和管理；
  * External网络：指VM网络之外的网络，不由neutron管理；
* Neutron架构：和其他OpenStack服务设计思路一样，也采用分布式架构，由多个组件(子服务)共同对外提供网络服务；
  * Neutron Server：对外提供OpenStack网络API，接收请求，并调用plugin处理请求；
  * Plugin：处理neutron Server发过来的请求，维护OpenStack逻辑网络状态，并调用agent处理请求；
  * Agent：处理Plugin的请求，负责在network provider上真正实现各种网络功能；
  * network provider：提供网络服务的虚拟或物理网络设备，例如Linux Bridge、Open vSwitch或者其他支持neutron的物理交换机；
  * Queue：neutron server、Plugin和Agent之间通过MQ通信和调用；
  * Database：存放OpenStack的网络状态信息，包括Network、Subnet、Port和Router等；
* Neutron Server组件详解
  * Core API：对外提供管理Network、Subnet、Port的RESTful API；
  * Extension API：对外提供管理Router、Load balance、firewall等资源的RESTful API；
  * common Service：认证和校验API请求；
  * neutron core：neutron server的核心处理程序，通过调用相应的Plugin处理请求；
  * core Plugin api：定义了core Plugin的抽象功能集合，neutron core通过该APi调用相应的core plugin
  * Extension plugin api：定义了Service plugin的抽象功能集合，neutron core通过该API调用相应的Service plugin；
  * core plugin：实现了core plugin api，在数据库中维护Network、Subnet和Port的状态，并负责调用相应的Agent在Network provider上执行相关操作，如创建Network
  * Service plugin：实现了Extension plugin api，在数据库中维护Router，LB、Security Group等资源的状态，并负责调用相应的Agent在network provider上执行相关操作，如创建router
  * 总结：Neutron Server包括两部分：提供API服务，运行Plugin，即 **Neutron Server = API + Plugin**

### [核心组件-Horizon管理界面](https://www.cnblogs.com/cloudhere/p/10812252.html)
* Horizon介绍： 为OpenStack提供一个Web前端的管理界面，通过Horizon所提供的Dashboard，管理员可以通过WebUI对OpenStack整体云环境进行管理，并观察各种操作结果与运行状态。

### [核心组件-Cinder存储服务](https://www.cnblogs.com/cloudhere/p/10815241.html)
* Cinder介绍
  * 理解Block Storage：操作系统获得存储空间的方式一般有两种：
    * 通过某种协议(SAS、SCSI、SAN、iSCSI等)挂接裸硬盘，然后分区、格式化、创建文件系统；或者直接使用裸硬盘存储数据(数据库)
    * 通过NFS、CIFS等协议，mount远程的文件系统
    * 第一种裸硬盘的方式叫Block Storage(块存储)，每个裸硬盘通常也被称作Volume(卷)，第二种叫做文件系统存储。NAS和NFS服务器，以及各种分布式文件系统提供的都是这种存储；
  * 理解Block Storage Service(Cinder)：Block storage service提供对volume从创建到删除整个生命周期的管理。从instance的角度看，挂载的每一个volume都是一块硬盘，Cinder的具体功能：
    * 提供REST API让用户能够查询和管理volume、volume snapshot以及vollume type
    * 提供scheduler调度volume创建请求，合理优化存储资源的分配
    * 提供Driver架构支持多种backend(后端)存储方式，包括LVM、NFS、Ceph和其他例如EMC、IBM等商业存储产品和方案；
* Cinder组件
  * cinder-api：接收api请求(和volume生命周期相关)，调用cinder-volume，是整个cinder的门户组件。所有cinder的请求都首先由cinder-api处理。cinder-api向外界暴露若干HTTP rest api接口。Keystone中可以查询对应endpoint。
  * cinder-volume：管理volume服务，与volume-provider协调工作，管理volume的生命周期。运行cinder-volume服务的节点被称为存储节点；
  * Driver架构：通过Driver架构支持多种volume-provider。cinder-volume为provider定义了统一的接口，provider只需要实现这些接口，就可以以Driver的形式即插即用到OpenStack系统中。
  * volume-provider：数据的存储设备，为 volume 提供物理存储空间。 cinder-volume 支持多种 volume provider，每种 volume provider 通过自己的 driver 与cinder-volume 协调工作。
  * volume-scheduler：通过调度算法选择最合适的存储节点创建volume。创建volume时，scheduler会基于容量、VolumeType等条件选择最合适的存储节点，然后创建volume。
  * MQ：子服务通过消息队列进行通信和相互协作。
  * Database Cinder：有些数据需要存放到DB中，一般使用MySQL。安装在控制节点中的，名称一般为`cinder`。
  * 物理部署方案：Cinder服务部署在控制节点和存储节点上，一般cinder-api和cinder-scheduler部署在控制节点中；cinder-volume部署在存储节点中；
    不过一个节点既可以是控制节点，也可以是存储节点，所以这些服务是有可能是部署在同一个节点中的。
    至于cinder-provider放在那里，provider是独立的，cinder-volume使用Driver和volume通信丙协调工作。所以只需要将Driver和cinder-volume放一起就可以了。
* Cinder设计思想：沿袭其他的服务，如Nova的设计思想。
  * API前端
    * cinder-api：作为Cinder组件对外的唯一窗口，向客户暴露Cinder能够提供的功能。
    * 优点：
      * 对外提供同一接口，隐藏实现细节
      * API提供REST标准调用服务，便于与第三方系统集成
      * 可以通过运行多个API服务实例实现高可用
  * Scheduler调度服务：Cinder可以有多个存储节点，当需要创建volume时，cinder-scheduler会根据存储节点的属性和资源使用情况选择一个最合适的节点来创建volume。
  * Worker工作服务：调度服务只管分配任务，cinder-volume作为Worker是真正执行任务的工作服务。Scheduler和Worker划分职能，使OpenStack非常容易横向扩展。
  * Driver框架：OpenStack作为开放的IaaS云操作系统，支持业界各种优秀的技术，这些技术可以是开源的，也可以是商业收费的。存储节点支持多种cinder-provider，只要实现统一的driver接口即可。

## 参考链接
* [云计算OpenStack](https://www.cnblogs.com/cloudhere/category/1439151.html)
* [整体介绍](https://jckling.github.io/2021/04/02/OpenStack/OpenStack%20%E6%95%B4%E4%BD%93%E4%BB%8B%E7%BB%8D/index.html)









