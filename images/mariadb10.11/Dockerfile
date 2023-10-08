FROM debian12.x

RUN apt install -y mariadb-server &&\
    echo "[mysqld]\nskip-networking=0\nskip-bind-address"  >> /etc/mysql/my.cnf

ADD init.txt /
# sleep时间太短,mariadb可能还没起来
RUN /usr/bin/mysqld_safe --nowatch &&\
    sleep 5 &&\
    mysql < /init.txt

# mysql -h221.mm -uzian -pzian
CMD /usr/bin/mysqld_safe --nowatch ; /usr/sbin/sshd -D