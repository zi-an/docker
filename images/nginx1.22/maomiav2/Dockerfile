FROM debian12.x

RUN apt install nginx -y

# 使用了proxy_pass+cache缓存加速
RUN sed -i "13 i proxy_cache_path /opt keys_zone=my_cache:10m;" /etc/nginx/nginx.conf &&\
    sed -i "14 i proxy_cache my_cache;proxy_cache_key \$request_uri;" /etc/nginx/nginx.conf &&\
    sed -i "15 i proxy_cache_valid 200 240h;" /etc/nginx/nginx.conf &&\
    sed -i "16 i proxy_hide_header Set-Cookie;" /etc/nginx/nginx.conf &&\
    sed -i "s|location /|location /default|g" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location =/ {rewrite / /index/home.html redirect;}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location / {proxy_pass https://maomi.com/;}" /etc/nginx/sites-enabled/default


ENV maomi=www.f89bb88d4062.com
ADD update.sh /
RUN chmod 700 /update.sh &&\
    cp /etc/nginx/sites-enabled/default /default &&\
    echo $maomi | awk -F '.' '{printf $2}' > url.txt &&\
    echo >> url.txt &&\
    sed -i "s|maomi.com|$maomi|g" /etc/nginx/sites-enabled/default

CMD nginx -g "daemon off;"
# 浏览器会因缓存原因,更新不及时
# 定时任务
# */5 * * * * docker exec -it maomiav2 /update.sh