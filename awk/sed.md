## sed格式

    - 命令行
      - sed [optiongs] 'command' files
    - 脚本格式
      - sed -f scriptfile files


#### 命令行模式
    - options 
      - -e
      - -n
    - command: 行定位的正则+ sed命令

```
1. sed -n '/root/p'
2. sed -e '10,20d' -e 's/flase/ture/g'
```


#### sed 基本命令
- -p :打印 一般和-n 配合使用
- 定位一行 ：x ; /pattern/
- 定位几行： x,y; /pattern/,x
- 取反行： x,y!
- 间隔航： first~step 


#### sed 操作命令
- -a 新增  
  - sed '5a -----------------'
- -i 插入
  - sed '5i -----------------'
- -c 替代
  - sed '5c -------------------'
- -d 删除
  - sed '5d '

```
#在files 文件尾部追加prot 5212 换行 port 21234
sed '$a prot 5212 \n port2 123556' files
# 打印error日志
sed -n '/error/p' files
#删除空行
sed '/^$/d' files
```

- -s 替换命令 ：分隔符、，#等
  - sed 's /false/ture/' passwd
- -g 整行选中
  - sed 's/flase/ture/g' passwd

```
ifconfi ens0 | sed 's/ine.*r://g'| sed 's/B.*//g'
```
### sed高级操作命令
- {}：多个sed命令 用；分开
  - sed {p;p;p} passwd
- -n :读取下一个输入行
  - sed {p;n;p} passwd
- & ：简化替换中的操作
  - sed 's/abc /&  /g' 
  -  