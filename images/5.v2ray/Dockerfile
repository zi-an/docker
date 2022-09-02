FROM centos7.8

RUN yum install unzip -y &&\
    wget https://github.com/v2fly/v2ray-core/releases/download/v4.45.2/v2ray-linux-64.zip &&\
    unzip v2ray-linux-64.zip &&\
    rm -rf /var/cache/yum/

# 配置文件除密码外已经加入,其中head命令会删除最后一个字符
ADD temp.json /
ENV TZ=Asia/Shanghai

RUN echo "export passwd=\`curl https://www.youhou8.com/pwd2\`" > update.sh &&\
    echo "sed \"s|123456|\$passwd|g\" /temp.json > config.json" >> update.sh &&\
    chmod +x update.sh &&\
    /update.sh

CMD ["/usr/sbin/sshd -D"]


# crontab
# 02 00 * * * docker exec v2ray sh "/update.sh"
# 03 00 * * * docker restart v2ray
