FROM debian12.x

# 开启容器后可能要十分钟后才能稳定同步给客户端,正常现象
RUN apt install -y chrony
RUN sed -i "s|2.debian.pool.ntp.org|ntp.aliyun.com|g" /etc/chrony/chrony.conf &&\
    echo "allow all" >> /etc/chrony/chrony.conf

CMD chronyd -d -x