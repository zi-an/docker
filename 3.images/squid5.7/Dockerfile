FROM debian12.x

RUN apt install -y squid &&\
    sed -i "s|http_access deny all|http_access allow all|g" /etc/squid/squid.conf
RUN sed -i "1irequest_header_access Via deny all" /etc/squid/squid.conf &&\
    sed -i "1irequest_header_access X-Forwarded-For deny all" /etc/squid/squid.conf &&\
    sed -i "1irequest_header_access X-Real-Ip deny all" /etc/squid/squid.conf
#RUN sed -i "1iaccess_log /logs/squid.log" /etc/squid/squid.conf
#/var/log/squid/access.log 使用宿主的挂载文件启动不了

CMD ["squid", "-N"]

# 日志配置 :http://www.squid-cache.org/Doc/config/logformat/
# 因性能原因,日志没有年月日

