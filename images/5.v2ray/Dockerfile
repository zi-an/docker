FROM centos7.8

RUN yum install unzip -y &&\
    wget https://github.com/v2fly/v2ray-core/releases/download/v4.45.2/v2ray-linux-64.zip &&\
    unzip v2ray-linux-64.zip &&\
    rm -rf /var/cache/yum/

# 配置文件除密码外已经加入,其中head命令会删除最后一个字符
ADD *.config.json /

RUN echo "curl https://youhou8.com/pwd2 > pwd" > update.sh &&\
    echo "head  -c -1 -q 1.config.json pwd 2.config.json > config.json" >> update.sh &&\
    chmod +x update.sh &&\
    /update.sh

CMD ["/usr/sbin/sshd -D"]


# crontab
# 02 00 * * * docker exec v2ray sh "/update.sh"
# 03 00 * * * docker restart v2ray
