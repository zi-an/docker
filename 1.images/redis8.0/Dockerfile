FROM debian13.x

RUN apt update &&\
    apt install -y redis-server

RUN sed -i 's|bind 127.0.0.1|bind 0.0.0.0|g' /etc/redis/redis.conf &&\
    sed -i 's|protected-mode yes|protected-mode no|g' /etc/redis/redis.conf &&\
    sed -i 's|# requirepass foobared|requirepass zian|g' /etc/redis/redis.conf &&\
    sed -i 's|daemonize yes|daemonize no|g' /etc/redis/redis.conf

RUN rm -rf /var/lib/apt &&\
    find /* -path /sys -prune -o -path /proc -prune -o -exec touch -ht 197001010000.00 {} \;

CMD ["redis-server", "/etc/redis/redis.conf"]