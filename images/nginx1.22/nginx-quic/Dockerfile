FROM debian12.1

# 建议使用火狐调试
RUN echo "deb https://nginx.org/packages/mainline/debian bookworm nginx" > /etc/apt/sources.list.d/nginx.list &&\
    apt update --allow-insecure-repositories &&\
    apt install -y nginx --allow-unauthenticated

# 创建10年私人证书,2048已经是最短的密钥
ADD cert.txt /
RUN openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout /etc/nginx/cert.key -out /etc/nginx/cert.pem < /cert.txt &&\
    wget https://nginx.org/favicon.ico -O /usr/share/nginx/html/favicon.ico

# idea2023有nginx配置文件的插件
# 如果开启http2,则火狐浏览器不会主动请求http3,也许自签名原因谷歌浏览器一直不使用http3协议
ADD quci.conf /etc/nginx/conf.d/
CMD ["/usr/sbin/sshd","-D"]

