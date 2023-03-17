
* 增删改查
* 创建视图、存储过程、函数
* 触发器、事件、事务、并发
* 数据库设计
* 索引以提高效率
* 保护数据库

# SELECT子句详解

```sql
USE sql_store; -- 选中需要使用的数据库
SELECT *
FROM customers
WHERE customer_id = 1
ORDER BY first_name
-- 双减号可以用来写注释
```

```sql
SELECT
last_name,
first_name,
points AS "POINTS HOHO", -- AS后跟别名，如果别名有空格，需要前后加引号
points * 10 + 100 AS money_they_obtain
FROM customers
```

```sql 
SELECT DISTINCT state -- DISTINCT关键字，不出现重复的
FROM customers;
```


# WHERE子句详解
对所有句子进行遍历

sql里
= 就是等于
<> 与 != 都是不等于
AND 是 并且
OR 是 或者
NOT 是 非

优先级 NOT > AND > OR

## IN关键字
`state = 1 OR state = 2 OR state = 3`  
等同于  
`state IN (1, 2, 3)`

## BETWEEN关键字
某个东西在值A与值B间（包括值A与B）

`WHERE points >= 0 AND points <= 100`  
等同于  
`WHERE points BETWEEN 0 AND 100`

## LIKE关键字
`WHERE name LIKE '%a%'; -- 这句也能搜到大写A`  
寻找name中有a的人

_代表一个字符，%代表任意多字符

## REGEXP关键字
`WHERE name REGEXP 'a';` 等同于`WHERE name LIKE '%a%';`

`WHERE name REGEXP '^abc';`  
表示name必须要abc开头

`WHERE name REGEXP 'abc$';`  
表示name必须要abc结尾

`WHERE name REGEXP 'a|b';`  
表示name必须要包含a或者b

`WHERE name REGEXP '[abc]d'; -- '[a-m]n' 表示从a到m都行，！！！不包括m`  
表示name必须要ad或bd或cd

## IS NULL关键字
```sql
WHERE phone IS NULL
WHERE phone IS NOT NULL
```

## ORDER BY
```sql
SELECT * , quantity * unit_price AS total_price
FROM order_items
WHERE order_id = 2
ORDER BY total_price; # 按照总价来排序
ORDER BY total_price DESC; # 按照总价来排序，但是是降序
```


## LIMIT
总是写在结尾
```sql
SELECT *
FROM customers
LIMIT 3; # 只显示三条
# LIMIT 6, 3; # 只显示三条，且前6条不显示，相当于偏移量为6
```