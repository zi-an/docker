# proxy加速

## Dockerfile是两种方案合一
``` 
proxy
├── /opt           : proxy_cache 缓存存储,内容更新
└── /var/www/html/ : proxy_store 永久存储,内容不变
```
## proxy_cache
```
proxy_cache_path /opt keys_zone=my_cache:10m; #定义分配缓存空间
proxy_cache my_cache; #使用my_cache的空间
proxy_cache_valid 200 1h; #状态200的记录缓存一小时
proxy_set_header Accept-Encoding ""; #防止使用gzip压缩,缓存出问题
```
## proxy_store
``` 
proxy_store on; #打开持久化存储
proxy_temp_path /var/www/html; #持久化存储目录,建议与root或者alias相同
if (!-e $request_filename){} #使用if判断文件是否存在,不存在时进行下一步操作
```

## 日志输入格式
* $request_time为请求处理时间（以秒为单位），分辨率为毫秒
* log_format的main是其ID,需要在access_log后面注明才会使用,否则都是默认日志
* time_iso8601比较标准,默认的time_local不好看
* access_log是已存在的,需要找到用
```
log_format main '$remote_addr [$time_iso8601] "$request" $status $request_time';
access_log logs/access.log main;
```