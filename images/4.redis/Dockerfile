FROM centos7.8

RUN echo "[qinghua-epel]" >> /etc/yum.repos.d/epel.repo &&\
    echo "name=qinghua-epel" >> /etc/yum.repos.d/epel.repo &&\
    echo "baseurl=https://mirrors.tuna.tsinghua.edu.cn/epel/7/x86_64/" >> /etc/yum.repos.d/epel.repo &&\
    echo "gpgcheck=0" >> /etc/yum.repos.d/epel.repo &&\
    yum install redis -y &&\
    rm -rf /var/cache/yum/

# 注释的是以守护进程运行
RUN  sed -i "s/bind 127.0.0.1/bind 0.0.0.0/" /etc/redis.conf &&\
     sed -i "s/protected-mode yes/protected-mode no/" /etc/redis.conf &&\
     #sed -i "s/daemonize no/daemonize yes/" /etc/redis.conf &&\
     sed -i "s/# requirepass foobared/requirepass haiyang/" /etc/redis.conf


CMD ["/usr/bin/redis-server","/etc/redis.conf"]