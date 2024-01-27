FROM debian12.x

# debian自带的nginx内的--with-stream=dynamic,是动态的,要重新编译,可能无法使用straem进行tcp转发
RUN apt install dnsmasq -y

CMD dnsmasq -dq
#dnsmasq --help