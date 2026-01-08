FROM debian13.x

# 设置环境变量，避免交互式提示
ENV DEBIAN_FRONTEND=noninteractive

RUN apt update &&\
    apt install -y postgresql &&\
    apt-get clean &&\
    mkdir -p /data &&\
    chown -R postgres:postgres /data &&\
    su - postgres -c "/usr/lib/postgresql/17/bin/initdb -D /data" &&\
    sed -i "s|#listen|listen_addresses = '*'#listen|g" /data/postgresql.conf &&\
    sed -i '$ahost all all 0.0.0.0/0 md5'  /data/pg_hba.conf

USER postgres
CMD ["/usr/lib/postgresql/17/bin/postgres", "-D", "/data"]
# sudo -u postgres psql -c "ALTER USER postgres WITH ENCRYPTED PASSWORD 'zian';"

