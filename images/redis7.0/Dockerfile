FROM debian12.x

RUN apt install -y redis-server
RUN sed -i "s/bind 127.0.0.1/bind 0.0.0.0/" /etc/redis/redis.conf &&\
    sed -i "s/protected-mode yes/protected-mode no/" /etc/redis/redis.conf &&\
    #sed -i "s/daemonize no/daemonize yes/" /etc/redis.conf &&\
    sed -i "s/# requirepass foobared/requirepass zian/" /etc/redis/redis.conf
