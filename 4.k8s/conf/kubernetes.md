# kubernetes 1.23.17单机版

# kubernetes介绍
kubectl 命令行工具
kubeadm 集群主机
kubelet 集群节点
根据官网,1.23是最后一个支持docker的版本参考

*https://blog.csdn.net/qq_15045983/article/details/127678302*

# 关闭swap+host+集群参数
```
swapoff -a
systemctl --type swap --all
systemctl mask dev-sda3.swap

hostname k610d
echo k610d > /etc/hostname
echo "192.168.10.5 k610d" >> /etc/hosts

mkdir /etc/docker
echo '{"hosts":["unix:///var/run/docker.sock","tcp://0.0.0.0:2375"],"registry-mirrors": ["http://hub-mirror.c.163.com"]}' > /etc/docker/daemon.json
echo "net.bridge.bridge-nf-call-ip6tables = 1" > /etc/sysctl.d/k8s.conf
echo "net.bridge.bridge-nf-call-iptables = 1" >> /etc/sysctl.d/k8s.conf
echo "net.ipv4.ip_forward = 1" >> /etc/sysctl.d/k8s.conf
```

# 安装kubeadm=1.23.17-00 kubelet=1.23.17-00 kubectl=1.23.17-00
* kubelet 安装完会自动启动
* kubectl completion bash完善代码补全
```
echo "deb http://mirrors.tuna.tsinghua.edu.cn/kubernetes/apt kubernetes-xenial main" > /etc/apt/sources.list.d/k8s.list
apt update --allow-insecure-repositories
apt install -y kubeadm=1.23.17-00 kubelet=1.23.17-00 kubectl=1.23.17-00 --allow-unauthenticated
kubectl completion bash >/etc/bash_completion.d/kubectl
```

# kubectl自动补全+运行配置(找不到/etc/kubernetes/admin.conf可以以后运行)
*不加--pod-network-cidr=10.244.0.0/16,kube-flannel.yml会一直CrashLoopBackOff重启
* kube-flannel.yml网络代理
* 默认master是pod的污点,不能代理,这里要忽略掉
```
kubeadm init --image-repository=registry.aliyuncs.com/google_containers --pod-network-cidr=10.244.0.0/16
mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
chown $(id -u):$(id -g) $HOME/.kube/config

wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
kubectl taint nodes --all node-role.kubernetes.io/master-
kubectl apply -f kube-flannel.yml 

#后台cni报错解决
sed -i "s|--network-plugin=cni| |g" /var/lib/kubelet/kubeadm-flags.env
systemctl restart kubelet.service
```
=================================================================================================================

# 面板
```
wget https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
kubectl apply -f recommended.yaml
kubectl proxy
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
*获取普通token的命令,不支持操作yaml等
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')
```
#查看节点日志
```
journalctl -u kubelet.service -f
```

#卸载k8s+docker
```
apt remove kubeadm kubelet kubectl --purge
apt remove docker.io --purge
rm /etc/kubernetes/ -rf
kubeadm reset
rm -rf /var/lib/etcd/
```

# 断电重启
* 任务多,有点慢
```
systemctl restart kubelet.service
docker rm $(docker ps -a -q)
``` 