#!/bin/sh

echo "update server"
unzip -o /home/linaro/Desktop/golang/tmp/fw.zip -d /home/linaro/Desktop/golang/ &

wait
echo "update end"
sync
echo "sync end"
