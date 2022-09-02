FROM centos7.8

RUN yum install unzip -y &&\
    wget https://github.com/XTLS/Xray-core/releases/download/v1.5.5/Xray-linux-64.zip &&\
    unzip Xray-linux-64.zip &&\
    rm -rf /var/cache/yum/ Xray-linux-64.zip

# 配置文件除密码外已经加入,其中head命令会删除最后一个字符
ADD *.json /

RUN echo "export passwd=\`curl https://www.youhou8.com/pwd2\`" > update.sh &&\
    echo "sed -i \"s|123456|\$passwd|g\" outbounds.json" >> update.sh &&\
    chmod +x update.sh &&\
    /update.sh

CMD ["/usr/sbin/sshd -D"]

# 相比v2ray,最好的是可以配置分开
