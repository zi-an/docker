FROM centos7.8

RUN yum install mariadb-server -y &&\
    rm -rf /var/cache/yum/

RUN mysql_install_db --user=mysql
ADD init.sql .

CMD ["/usr/bin/mysqld_safe"]
# 打开容器后需要执行: mysql < init.sql