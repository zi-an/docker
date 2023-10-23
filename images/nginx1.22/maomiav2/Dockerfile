FROM debian12.x

RUN apt install nginx -y

# 使用了proxy_pass+cache缓存加速,配置和实际大小不符,导致无法缓存
RUN sed -i '13 i proxy_cache_path /opt keys_zone=my_cache:10m;proxy_cache my_cache;proxy_cache_valid 200 1h;proxy_set_header Accept-Encoding "";' /etc/nginx/nginx.conf

RUN sed -i 's|location /|location /default|g' /etc/nginx/sites-enabled/default &&\
    sed -i '42 i location ~ .detail. {proxy_store /var/www/html$request_uri;if (!-e $request_filename){proxy_pass https://www.maomi.com;}}' /etc/nginx/sites-enabled/default &&\
    sed -i '42 i location / {proxy_pass https://www.maomi.com;proxy_cache_key $request_uri;}' /etc/nginx/sites-enabled/default &&\
    sed -i '42 i location =/ {rewrite / /index/home.html redirect;}' /etc/nginx/sites-enabled/default

ENV maomi=www.f89bb88d4062.com
ADD update.sh /
RUN chmod 700 /update.sh &&\
    cp /etc/nginx/sites-enabled/default /default &&\
    echo $maomi | awk -F '.' '{printf $2}' > url.txt &&\
    echo >> url.txt &&\
    sed -i "s|www.maomi.com|www.$maomi.com|g" /etc/nginx/sites-enabled/default

CMD nginx -g "daemon off;"
# 定时任务,五分钟一次,docker exec -it 需要创建控制台,这里不适用 -it 直接执行
# */5 * * * * docker exec maomiav2 /update.sh