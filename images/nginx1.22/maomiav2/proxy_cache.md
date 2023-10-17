# cache加速

# 方案一,永久不变
* 使用proxy_store永久缓存
``` 
sed -i "42 i location =/ {rewrite / /index/home.html redirect;}" /etc/nginx/sites-enabled/default &&\
sed -i "42 i location / {proxy_store on;proxy_temp_path /www/;alias /www/;if(!-e $request_filename){proxy_pass https://www.maomi.com;}}" /etc/nginx/sites-enabled/default
```

# 方案二,会变的方案
```
sed -i "13 i proxy_cache_path /opt keys_zone=my_cache:10m;" /etc/nginx/nginx.conf &&\
sed -i "14 i proxy_cache my_cache;proxy_cache_key \$request_uri;" /etc/nginx/nginx.conf &&\
sed -i "15 i proxy_cache_valid 200 240h;" /etc/nginx/nginx.conf &&\
sed -i "16 i proxy_hide_header Set-Cookie;" /etc/nginx/nginx.conf &&\
sed -i "s|location /|location /default|g" /etc/nginx/sites-enabled/default &&\
sed -i "42 i location =/ {rewrite / /index/home.html redirect;}" /etc/nginx/sites-enabled/default &&\
sed -i "42 i location / {proxy_pass https://maomi.com/;}" /etc/nginx/sites-enabled/default
```

detail->使用永久缓存
list->使用1小时