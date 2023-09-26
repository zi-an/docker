# wireshark + tcpdump

## 安装+打开sshd的x11
* wireshark 分析软件
* tcpdump 抓取软件
* fonts-wqy-zenhei 中文字体
```
apt install -y wireshark tcpdump fonts-wqy-zenhei
echo "X11UseLocalhost no" >> /etc/ssh/sshd_config
```

## 测试
* 准备
```
apt install -y nginx
nginx
```

* 抓取查看
* 抓取过程中,需要在另一台机器上使用curl命令
```
tcpdump port 80 -w 80.cap
wireshark -r 80.cap
```