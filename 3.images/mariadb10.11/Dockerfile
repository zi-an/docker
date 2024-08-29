FROM debian12.x

# 安装mariadb开启远程
RUN apt install -y mariadb-server &&\
    echo "[mysqld]\nskip-networking=0\nskip-bind-address"  >> /etc/mysql/my.cnf

ADD init.txt /
ADD mariadb.txt /mariadb.sh
# sleep时间太短,mariadb可能还没起来
RUN rm -rf /var/lib/mysql &&\
    chmod 700 /mariadb.sh

# mysql -h221.mm -uzian -pzian
CMD /mariadb.sh ; /usr/sbin/sshd -D