面向对象的程序设计中，对象的关系主要有：
composition，组合。比较强的包含关系。例如一辆车拥有四个轮子。
aggregation，聚合，不稳定的包含关系。可两个对象单独存在，例如一个人跟它的地址。
association，联系。引用关系，一般比较长期，例如一个人跟他的订单。
继承。最紧密的关系。
---
类之间的关系
dependence 依赖 uses a
aggregation 聚合 has a
inheritance 继承 is a

--- 
在使用 goimports 命令后，git add 时，有如下提示  

```
warning: in the working copy of 'Go/Go_1000-/1475.go', LF will be replaced by CRLF the next time Git touches it
```

--- 
杰克的oop理解

三两句话说不清楚。 核心就是分离访问层、业务层、表现层。 访问层：模板数据、数据库、设备设置、缓存库等 表现层：UI、渲染的模型 业务层就是把访问层和表现层耦合在一起。 也就是访问层和表现层不处理业务逻辑，只处理自身的逻辑。业务层处理耦合。 说起来都是老生常谈的内容，三层架构或者MVC之类的。三年前也去这样设计就是掌握不了精髓，现在悟到了。

---

LRU算法
使用map存 < int , node>
node连成双向链表，每次 get 或者 put 都把对应的 node 放到链表头
node的内容，int key, int value, node *left, node *right

--- 
```
# 开头表示注释
以斜杠/开头表示目录；
以星号*通配多个字符；
以问号?通配单个字符
以方括号[]包含单个字符的匹配列表；
以叹号!表示不忽略(跟踪)匹配到的文件或目录；
```

---

Knuth洗牌算法
将需要洗牌的顺序存在一个数组num中，设数组尺寸为n
如此反复进行n轮，第i轮的操作：
将 num[i] 与其之后的某个数 num[j] 交换，其中 j 的取值范围为 [ i , n )
j 的生成方法 : j = i + rand() % (n - i)
如此便可保证生成的序列任何可能性都概率相等
#算法

--- 
我的HTTP请求头

GET / HTTP/1.1
Host: 106.14.0.56
Upgrade-Insecure-Requests: 1
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.2 Safari/605.1.15
Accept-Language: zh-cn
Accept-Encoding: gzip, deflate
Connection: keep-alive

---
windows
查看电池损耗情况的方法:
以管理员身份打开命令提示符(cmd），也就是一个黑框框
输入powercfg /batteryreport /output "C:\battery_report.html"
在c盘中找到生成的电池报告文件

---

# r9p Part


# m2015 Part