# ffmpeg 文档

| https://ffmpeg.org/ffmpeg-all.html

## 测试脚本
* 下载-->剪辑-->合并
```shell
wget https://heishenhua.com/video/b1/gamesci_wukong.mp4 -O hsh.mp4
ffmpeg -i hsh.mp4 -ss 00:07 -to 01:21 -c copy hsh1.mp4
ffmpeg -i hsh.mp4 -ss 01:34 -to 12:12 -c copy hsh2.mp4
echo "file hsh1.mp4" >> hsh.txt
echo "file hsh2.mp4" >> hsh.txt
ffmpeg -f concat -i hsh.txt -c copy heishenhua.mp4
``` 

## 查看视频信息
```shell
ffmpeg -i heishenhua.mp4
```

## 视频转码率 mp4转avi
```shell
ffmpeg -i bh3.mp4 bh3.avi
```
* 自动根据后缀转格式

## mp4转hls(m3u8)
```shell
ffmpeg -i heishenhua.mp4 -hls_time 5 -hls_list_size 0 -c copy heishenhua.m3u8
```
* -hls_time 每个切片时间5秒
* -hls_list_size 切片列表的数量,在m3u8文件里的记录数,0代表所有
* -c codec(编码)的简写,copy表示不转码

## rtmp推送(直播)
```shell

```

## 处理器/显卡相关
-c:v libx264 原生的264编码,纯靠CPU
-c:v h264_qsv 英特尔硬件加速
-c:v h264_amf AMD硬件加速
-c:v h264_nvenc 英伟达硬件加速
-c:v h264_videotoolbox MAC笔记本加速
* 速度快很多,但docker内不一定能用/加--gpus参数
```shell
ffmpeg.exe -i heishenhua.mp4 -c:v h264_amf -r 30 AMD30.mp4
```
原生264成绩6:15
5600处理器h264_amf成绩2:00