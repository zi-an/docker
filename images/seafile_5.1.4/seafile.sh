#!/bin/sh

if [ -d "/data" ] && [ ! -d "/data/seafile-server-5.1.4" ];
then mv /seafile/* /data/;
sed -i '3cif [ ! -e "/data/seafile-data/seafile.db" ];then'  /data/seafile-server-5.1.4/seafile.sh;
sed -i '4c/data/seafile-server-5.1.4/setup-seafile.sh</answer && sleep 10;fi;'  /data/seafile-server-5.1.4/seafile.sh;
fi

/data/seafile-server-5.1.4/seafile.sh start
/data/seafile-server-5.1.4/seahub.sh start 80
/bin/sh
