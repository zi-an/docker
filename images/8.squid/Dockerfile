FROM centos7.8

# 免用户验证,隐藏代理信息
RUN yum install squid httpd-tools -y &&\
    echo "request_header_access Via deny all" >> /etc/squid/squid.conf &&\
    echo "request_header_access X-Forwarded-For deny all" >> /etc/squid/squid.conf

# 以下是删除默认配置，并配置验证用户48小时有效
RUN htpasswd -bc /etc/squid/passwd zian zian &&\
    sed -i "s|acl |#acl |g" /etc/squid/squid.conf &&\
    sed -i "s|http_access |#http_access |g" /etc/squid/squid.conf &&\
    echo "auth_param basic credentialsttl 48 hours" >> /etc/squid/squid.conf &&\
    echo "auth_param basic program /usr/lib64/squid/basic_ncsa_auth /etc/squid/passwd" >> /etc/squid/squid.conf &&\
    echo "acl auth_user proxy_auth zian" >> /etc/squid/squid.conf &&\
    echo "http_access allow auth_user" >> /etc/squid/squid.conf

CMD ["/usr/sbin/sshd -D"]