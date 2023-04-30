FROM centos7.8

# nginx在epel库内
RUN wget https://mirrors.tuna.tsinghua.edu.cn/centos-vault/7.8.2003/extras/x86_64/Packages/epel-release-7-11.noarch.rpm --no-check-certificate &&\
    rpm -i epel-release-7-11.noarch.rpm &&\
    yum install nginx iproute -y

# 增加nginx代理规则
RUN sed -i "50 i location / {proxy_pass http://192.168.10.210;}" /etc/nginx/nginx.conf

CMD ["/usr/sbin/sshd -D"]