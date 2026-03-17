FROM debian13.x

# 安装mariadb开启远程,开启binlog,设置binlog保留30天,binlog大小1G
RUN apt update &&\
    apt install -y mariadb-server

RUN mkdir /var/log/mariadb &&\
    echo "[mysqld]\nskip-networking=0\nskip-bind-address"  >> /etc/mysql/my.cnf &&\
    echo "log_bin=/var/lib/mysql/mariadb-bin\nserver_id=1\nbinlog_format = ROW" >> /etc/mysql/my.cnf &&\
    echo "expire_logs_days=0\nmax_binlog_size=1073741824" >> /etc/mysql/my.cnf

ADD init.txt mariadb.sh /

# sleep时间太短,mariadb可能还没起来
RUN rm -rf /var/lib/mysql

RUN rm -rf /var/lib/apt &&\
    find /* -path /sys -prune -o -path /proc -prune -o -exec touch -ht 197001010000.00 {} \;

# mysql -h221.mm -uzian -pzian
CMD /mariadb.sh ; /usr/sbin/sshd -D