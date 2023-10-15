FROM debian12.x

RUN apt install nginx -y

# 使用了proxy_pass+cache缓存加速,配置和实际大小不符,导致无法缓存
RUN sed -i "s|location /|location /default|g" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location =/ {rewrite / /index/home.html redirect;}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location / {if (!-e \$request_filename){proxy_pass https://www.maomi.com;}}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i proxy_store on;proxy_temp_path /var/www/html/;proxy_set_header Accept-Encoding '';" /etc/nginx/sites-enabled/default

ENV maomi=www.f89bb88d4062.com
ADD update.sh /
RUN chmod 700 /update.sh &&\
    cp /etc/nginx/sites-enabled/default /default &&\
    echo $maomi | awk -F '.' '{printf $2}' > url.txt &&\
    echo >> url.txt &&\
    sed -i "s|www.maomi.com|www.$maomi.com|g" /etc/nginx/sites-enabled/default

CMD nginx -g "daemon off;"
# 浏览器会因缓存原因,更新不及时
# 定时任务
# */5 * * * * docker exec -it maomiav2 /update.sh