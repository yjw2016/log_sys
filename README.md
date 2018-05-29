1.项目背景  
- a. 收集日志  
- b.机器比较少时，登录服务器产看即可  
- c.集群规模巨大，登录机器查看不满足工程需求

2.方案
- 收集日志，统一存储
- 建立索引，实现日志搜索
- 提供web界面，访问日志搜索
3.需求
- 数据量大
- 低延迟
- 可扩展

4.已有方案
- ELK

5.又有方案弊端
- 配置量大
- 无法准确获取logstash状态
- 无法定制化

6.HADOOP ES STORM => KAFKA =>LOG AGENT

7.组件
- Log Agent  做手机
- kafka     做异步，解耦
- zk        服务注册&发现
- ES,elasticsearch 搜索
- Hadoop      分布式

8.安装kafka：
- 安装JDK
- 安装zookeeper
- 安装kafka

9.log agent
- kafka config tailf log

10.config -> n* tailf -> kafka

11.kafka示例代码

12.tailf组件使用

13.配置文件库使用 beego config

14.日志库的使用 beego logs










