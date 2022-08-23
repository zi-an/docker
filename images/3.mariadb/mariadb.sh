#!/bin/sh

if [ ! -e "/var/lib/mysql/mysql/" ] ; then
mysql_install_db --user=mysql;
/usr/bin/mysqld_safe --nowatch ;
sleep 3;
mysql < init.sql;
sh;
fi

/usr/bin/mysqld_safe

