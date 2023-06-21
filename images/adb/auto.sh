#!/bin/bash

date > ~/start.txt
adb connect 192.168.10.12

sleep 10s

for((;;))
do 
sleep 2
let x=2000+$[$RANDOM%10]
let y=1000+$[$RANDOM%10]
adb shell input tap $x $y;
#echo $x $y
done 

