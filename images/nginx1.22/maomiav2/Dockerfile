FROM debian12.x

RUN apt install nginx -y

# 使用了proxy_pass+cache缓存加速
RUN sed -i "13 i proxy_cache_path /opt keys_zone=my_cache:10m;" /etc/nginx/nginx.conf &&\
    sed -i "14 i proxy_cache_key \$request_uri;proxy_cache my_cache;" /etc/nginx/nginx.conf &&\
    sed -i "s|location /|location /default|g" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location =/ {rewrite / /index/home.html redirect;}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location / {proxy_pass https://maomi.com/;}" /etc/nginx/sites-enabled/default


ENV maomi=www.f89bb88d4062.com
RUN sed -i "s|maomi.com|$maomi|g" /etc/nginx/sites-enabled/default
CMD nginx -g "daemon off;"
# 以后要写脚本获取域名