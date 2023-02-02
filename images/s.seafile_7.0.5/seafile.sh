#!/bin/sh

if [ -d "/data" ] && [ ! -d "/data/seafile-server-7.0.5" ];
then mv /seafile/* /data/;
sed -i '3cif [ ! -e "/data/seafile-data/seafile.db" ];then'  /data/seafile-server-7.0.5/seafile.sh;
sed -i '4c/data/seafile-server-7.0.5/setup-seafile.sh</answer && sleep 10;fi;'  /data/seafile-server-7.0.5/seafile.sh;
fi

/data/seafile-server-7.0.5/seafile.sh start
/data/seafile-server-7.0.5/seahub.sh start
/bin/sh
