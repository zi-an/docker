#
#
```
wget 7.mm/json/demo.yaml
wget 7.mm/json/demo-server.yaml

kubectl apply -f demo.yaml 
kubectl apply -f demo-server.yaml
```

# 负载均衡测试
* 三台机分别执行
```
kubectl exec -it nginx-deployment#ID  sh
echo 1 >> /usr/share/nginx/html/t
echo 2 >> /usr/share/nginx/html/t
echo 3 >> /usr/share/nginx/html/t
```

* curl看输出结果的负载
```
curl 5.mm:30080/t
```
