# sed

sed 又称为流处理编辑器，和vim 不同的是sed是针对单行进行处理的。从文本或者管道获取到
输入内容，读取一行到模式空间（可理解为buffer），使用sed命令进行处理，输出到屏幕，然后循环
读取下一行内容到模式空间，重复上述内容。sed 本质是一次处理一行数据，然后输出，sed 不改变
源文件内容，如果使用 -i命令则是将源文件内容给替换掉了，使用-i命令一定要谨慎。
### 目录
 - sed [optiongs] 'command' files
 - [optios](####options)
 - [commands](####commands)    
    - [定位](####定位)
    - [命令](####命令)
    - [注意点](####注意点)
    
### sed options
命令行模式下一般通用的两个options ，-e: 选中script脚本来处理这个files，-n 是代表安静模式，
一般配合操作命令p来进行使用。一般其他的选项是用不到的。
  - -e ：选中特定的脚本
  - -n ：安静模式

````
-e<script>或--expression=<script> 以选项中指定的script来处理输入的文本文件。
-f<script文件>或--file=<script文件> 以选项中指定的script文件来处理输入的文本文件。
-h或--help 显示帮助。
-n或--quiet或--silent 仅显示script处理后的结果。
-V或--version 显示版本信息。
````
### sed commands
- commands = 定位操作+sed命令操作
````
sed '5d' files : 删除files 的第五行
sed '5a abbc' files : 在files第五行插入abbc
````
#### sed 定位操作
- 定位一行 ：x ; /pattern/
- 定位几行： x,y; /pattern/,x
- 取反行： x,y!
- 间隔行： first~step 
#### sed 操作命令
这里强调一点：command: 行定位的正则+ sed命令，但是s命令是特殊，需要特殊记忆。比如说我要删除
第10行： sed '10d' files,可以看到我们直接选中10行进行删除。 

- p 打印： 一般配合选项 -n 使用
- a 新增  
- i 插入
- c 替代
- d 删除
- s 替换 ：一般替换选中行第一个匹配对象，需要全部匹配要g 参数
````
sed '5a -----------------'
sed '5i -----------------'
sed '5c -------------------'
sed '5d ' ： 删除第五行
sed 's/abc/123/g' : 将abc替换成123
ifconfi ens0 | sed 's/ine.*r://g'| sed 's/B.*//g
````

#### 注意点
-  mac下面几乎aci操作都是会报错 
   ````
   command i expects \ followed by text
   ````
- mac 下面定位操作间隔行是一个不能被认识的指令
   ````
   sed: 1: "1~2p": invalid command code ~
   ````
- c 命令和s命令的区别在于c命令会把选中的行全部替换成我们新赋值的字符，而s
是针对行内的字符进行替换
### sed高级操作命令
- {}：多个sed命令 用；分开
- n :读取下一个输入行
- & ：简化替换中的操作
- () :更加简化&的操作
- rw: r 复制制定文件插入到匹配行，w 复制
- q: 退出sed 

````
sed {p;p;p} passwd
sed {p;n;p} passwd
sed 's/abc /&  /g' 
sed 's/iam\(good\)/\1/g' = goog 代替iamgood
````