#

# m3u8download脚本
* 授权后放到/usr/bin/下
```
#m3u8downloader
#!/bin/bash
default="-vcodec copy -acodec copy -hls_list_size 0 -hls_segment_filename .$2%03d.ts "
nohup ffmpeg -i $1 $default $2.m3u8 &
```