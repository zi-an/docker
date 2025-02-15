# debian docker镜像

## 下载到本地
* https://github.com/debuerreotype/docker-debian-artifacts/tree/dist-amd64/
* 这是powershell中的wget,与linux的不同
* -Uri与-OutFile必须指定才能下载
* 最好在笔记本上使用代理加速下载
```shell
wget.exe https://github.com/debuerreotype/docker-debian-artifacts/raw/refs/heads/dist-amd64/bookworm/oci/blobs/rootfs.tar.gz
```

