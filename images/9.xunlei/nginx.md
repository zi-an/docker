
# 安装软件
```
docker exec -it  mynas bash
apt update
apt install nginx -y
```
# 配置文件
* 因为没有vi,所以用cat+sed
```
cat > /etc/nginx/nginx.conf<<EOF
user www-data;
worker_processes  auto;

error_log  /var/log/nginx/error.log notice;
pid        /var/run/nginx.pid;


events {
    worker_connections  1024;
}

http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  65;

    #gzip  on;

    server {
    listen       80;
    server_name  localhost;
 
    location / {
        autoindex on;
        autoindex_exact_size off;
        autoindex_localtime on;
		charset utf-8;
        root /xunlei/downloads/;
    }
    }

}
EOF
```

# 编写跳转页
```
cat > /home/nginx/xunlei/xunlei.html<<EOF
<html>
<head><meta http-equiv="refresh" content="0;url=http://209.mm:2345"></head>
</html>
EOF
```