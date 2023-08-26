FROM guacamole/guacd:1.5.3
#debian12.1
#docker pull guacamole/guacd:1.5.3

# https://github.com/dushixiang/next-terminal/releases
# https://github.com/dushixiang/next-terminal/releases/download/v1.3.9/next-terminal.tar.gz
RUN wget https://github.com/dushixiang/next-terminal/releases/download/v1.3.9/next-terminal.tar.gz
RUN tar -zxvf next-terminal.tar.gz
#  /next-terminal/next-terminal
#  /opt/guacamole/sbin/guacd -b 0.0.0.0 -L $GUACD_LOG_LEVEL -f



CMD  /opt/guacamole/sbin/guacd -b 0.0.0.0 -L $GUACD_LOG_LEVEL -f;/next-terminal/next-terminal
#https://apache.org/dyn/closer.lua/guacamole/1.5.3/source/guacamole-server-1.5.3.tar.gz?action=download
#https://apache.org/dyn/closer.lua/guacamole/1.5.3/source/guacamole-client-1.5.3.tar.gz?action=download

#guacamole/guacd
#docker pull guacamole/guacamole:1.5.2