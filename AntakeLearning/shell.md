> 查看系统支持的shell

```shell
cat /etc/shells
[root@iZwz9almo8p830btq7voo9Z shellLearning]#  cat /etc/shells
/bin/sh
/bin/bash
/usr/bin/sh
/usr/bin/bash
```

> 查看系统当前的shell

```shell
[root@iZwz9almo8p830btq7voo9Z shellLearning]# echo $SHELL
/bin/bash
```

## Shell脚本文件编写规范

### 脚本文件后缀名规范

shell脚本文件就是一个文本文件，建议使用.sh结尾

### 首行格式规范

```shell
#!/bin/bash
```

> 含义：设置当前shell脚本采用bash解析器运行脚本代码

### 注释格式

```shell
# 单行注释
:<<!
# 多行注释内容
!
```

## 脚本文件常用执行的三种方式

1. sh解析器执行

   ```shell
   sh helloworld.sh
   ```

2. bash解析器执行

   ```shell
   bash helloworld.sh
   ```

3. 仅路径执行

   ```shell
   ./helloworld.sh
   #需要具有可执行权限
   [root@iZwz9almo8p830btq7voo9Z shellLearning]# chmod a+x helloworld.sh 
   [root@iZwz9almo8p830btq7voo9Z shellLearning]# ./helloworld.sh 
   hello world
   ```

   > 三种方式运行区别

   bash或sh执行是直接通过解析器执行，不需要可执行权限

   通过路径执行需要可执行权限

## Shell脚本入门：多命令处理

### 案例需求

已知目录  /root/shellLearning目录，执行bashch.sh脚本，在目录下创建一个one.txt,在one.txt中写入"hello shell"

### 步骤分析

1. 使用mkdir创建/root/shellLearning目录
2. 创建脚本文件batch.sh
3. 编辑脚本文件
   1. 创建文件one.txt
   2. 写入hello shell
4. 执行脚本文件

```shell
[root@iZwz9almo8p830btq7voo9Z shellLearning]# bash -v bashch.sh 
#!/bin/bash
touch /root/shellLearning/one.txt
echo "hello shell" >> /root/shellLearning/one.txt

[root@iZwz9almo8p830btq7voo9Z shellLearning]# bash bashch.sh 
[root@iZwz9almo8p830btq7voo9Z shellLearning]# ls
a.sh  a.txt  bashch.sh  hello.txt  helloworld.sh  one.txt  readme.txt  weizhi.sh  while.sh
[root@iZwz9almo8p830btq7voo9Z shellLearning]# cat one.txt 
hello shell
```

## Shell变量：环境变量

### 变量类型

1. 系统环境变量
2. 自定义变量
3. 特殊符号变量

#### 系统环境变量

##### 介绍

是系统提供的共享变量，是linux系统加载shell的配置文件中定义的变量共享给所有的shell程序使用

#### 自定义变量

#### 特殊符号变量

### Shell的配置文件分类

1. 全局配置文件

   /etc/profile

   /etc/profile.d/*.sh

   /etc/bashrc

2. 个人配置文件

   当前用户/.bash_profile

   当前用户/.bashrc

### 环境变量分类

在linux系统中，环境变量按照其作用范围不同大致可以分为系统级环境变量喝用户级环境变量

系统级环境变量：Shell环境加载全部配置文件中的变量共享给所有用户Shell程序使用，全局共享

用户级环境变量：Shell环境加载个人配置文件中的变量共享给当前用户Shell程序使用，登录用户使用

### 查看当前shell系统环境变量

> env

### 查看Shell变量（系统环境变量+自定义变量+函数）

> set

### 环境变量演示查看

```shell
[root@iZwz9almo8p830btq7voo9Z shellLearning]# echo $HOSTTYPE
x86_64
```

## Shell变量：自定义变量

定义在一个脚本文件中的变量，只能在该脚本文件中使用，就是局部变量。

> 定义语法

```shell
var_name=value
```

> 变量定义规则

1. 变量名称可以有字母，数字和下划线，但是不能以数字开头
2. 等号两侧不能有空格
3. 在bash环境中，变量的默认类型都是字符串类型，无法直接进行数值运算
4. 变量如果有空格，必须使用双引号括起来
5. 不能使用shell的关键字作为变量名称

> 查询变量语法

```shell
#$var_name
#${var_name}
#区别花括号方法适合拼接字符串
```

> 变量删除

```shell
#unset var_name
[root@iZwz9almo8p830btq7voo9Z shellLearning]# unset age
[root@iZwz9almo8p830btq7voo9Z shellLearning]# echo 我的名字是$name,年龄是$age
我的名字是zhangsan,年龄是
```

### 自定义常量

> 介绍

就是变量设置值以后不可以修改的变量叫常量，也叫只读变量

> 顶一规则

```shell
#readonly var_name
readonly sex=男
```

## 自定义全局变量

### 父子Shell环境介绍

例如：有2个脚本文件a.sh和b.sh

如果在a.sh中执行了b.sh脚本文件，那么a.sh就是父shell环境，b.sh就是子shell环境

### 自定义全局变量介绍

> 就是在当前脚本文件中定义了全局变量，这个全局变量可以在当前shell环境与子shell环境中使用

### 自定义全局变量语法

```shell
export var_name1 var_name2
```

> 例子

```shell
[root@iZwz9almo8p830btq7voo9Z shellLearning]# touch a.sh b.sh
[root@iZwz9almo8p830btq7voo9Z shellLearning]# ls
a.sh  a.txt  bashch.sh  b.sh  hello.txt  helloworld.sh  one.txt  readme.txt  weizhi.sh  while.sh
[root@iZwz9almo8p830btq7voo9Z shellLearning]# vim a.sh
[root@iZwz9almo8p830btq7voo9Z shellLearning]# vim b.sh
[root@iZwz9almo8p830btq7voo9Z shellLearning]# bash a.sh
b.sh中输出a.sh中的变量var4,值为
#修改为全局变量
[root@iZwz9almo8p830btq7voo9Z shellLearning]# bash a.sh
b.sh中输出a.sh中的变量var4,值为wocao
```

## Shell变量：特殊变量

> $n

```shell
$n用于接收脚本的参数
$0是脚本名称
$1-$9是1-9个参数
10+用${n}
```

> $#

```shell
所有输入参数的个数
```

> $*、$@

```shell
都是获取输入的所有参数
1.不使用双括号括起来，功能一样
2.使用双引号括起来
"$*"获取所有的参数拼接为一个字符串
"$@"获取以组参数列表对象
```

> 循环语法

```shell
for var in 列表变量
do
	语句
done
```

> $?上一句shell的执行状态码，0成功，非0失败

> $$ 用于获取当前的shell的进程id号

## shell系统环境变量深入

### 创建系统环境变量

1. 编辑/etc/profile全局配置文件

   ```shell
   #增加命令：定义变量VAR1=VAR1 并导出为环境变量
   ```

2. 重载配置文件/etc/profile,因为配置文件修改后需要立刻加载里面的数据

   ```shell
   source /etc/profile
   ```

3. 在shell环境中读取系统级环境变量VAR1

