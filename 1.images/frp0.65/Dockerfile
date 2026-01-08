FROM scratch

# https://github.com/fatedier/frp/releases/download/v0.65.0/frp_0.65.0_linux_amd64.tar.gz
ADD frp_0.65.0_linux_amd64.tar.gz /
ADD frpc.ini /
CMD ["/frp_0.65.0_linux_amd64/frpc"]

# vi frpc.ini
# docker cp frpc.ini frp:/ & docker restart frp
# 访问地址 public.freefrp.org:12121
