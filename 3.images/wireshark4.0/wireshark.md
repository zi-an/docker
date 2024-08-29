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

* 抓取查看80端口
* 抓取过程中,需要在另一台机器上使用curl命令
```
tcpdump port 80 -w 80.cap
wireshark -r 80.cap
```

### 流量统计
* 一般配合squid代理使用,过滤掉客户端IP
统计-->IPv4 Statistics--> All Addresses-->显示过滤器(ip.addr != 192.168.10.15)

### windows使用tcpdump
| https://www.microolap.com/products/network/tcpdump/
* cmd中运行需要以管理员权限

获取网卡信息+使用第2张网卡
```
tcpdump -D
tcpdump -i 2
tcpdump -i 2 -w z:\tcp.cap
```

tcpdump -D信息,第2张网卡是无线网卡
```
1.\Device\{F5DFA10B-7FE3-4DD8-83D5-5B5111133EB9} (Microsoft Wi-Fi Direct Virtual Adapter)
2.\Device\{F80ADA12-46F1-4204-A42B-4A7BA966C1F7} (Realtek 8852AE Wireless LAN WiFi 6 PCI-E NIC)
3.\Device\{0634F8AF-9504-4EC9-A5A3-C5EB9D742A8D} (Microsoft Wi-Fi Direct Virtual Adapter)
4.\Device\{D679DEBE-5B0E-44FE-B6AC-BC9977F93085} (Hyper-V Virtual Ethernet Adapter)
5.\Device\{D4909C24-0905-49FE-BC39-9E5990ECA7FC} (WAN Miniport (Network Monitor))
```