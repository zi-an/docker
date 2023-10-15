#!/bin/bash

url=`tail -n1 /url.txt`
url=`curl https://www.$url.com/ -o /dev/null -sw %{redirect_url}`
url=`echo $url|awk -F '.' '{printf $2}'`
if [ ${#url} -gt 6 ] ; then
  sed "s|www.maomi.com|www.$url.com|g" /default > /etc/nginx/sites-enabled/default
  date >> /url.txt
  echo $url >> /url.txt
  nginx -s reload
  echo 'update is ok'
fi