# 网络监测 nload+iftop

# 磁盘监测 lsof

# 文件夹大小du
* 文件深度为1的
```shell
du -hd1
```


## 查找改动文件
```shell
find / -mmin n/+n/-n
find / -mmin -3 ! -path "/sys/*" ! -path "/proc/*"
```
*  n 专指第n分钟
* +n 指第n分钟之前
* -n 指第n分钟之内
* ! -path 表示忽略后续文件,字符串要用""包含

## 树形目录
```shell
tree -ahL 2  /etc
```
* -a 显示所有文件,包括隐藏文件
* -h 已合适单位显示文件大小
* -L 2展开二级目录




