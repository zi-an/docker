#!/bin/sh

if [ ! -e "/var/lib/mysql/mysql/" ];
then
	mariadb-install-db --user=mysql
	/usr/bin/mariadbd-safe --nowatch
	sleep 3
	mariadb <init.txt
	sh
else
	/usr/bin/mariadbd-safe --nowatch
fi
