#!/bin/sh

echo "update server"
unzip -o /home/linaro/Desktop/golang/tmp/fw.zip -d /home/linaro/Desktop/golang/ &

wait
echo "update end"
sync
echo "sync end"
chmod 777 /home/linaro/Desktop/
chmod 777 /home/linaro/Desktop/golang/
chmod 777 /home/linaro/Desktop/golang/www
chmod 777 /home/linaro/Desktop/golang/www/html/
chmod 777 /home/linaro/Desktop/golang/www/js/
chmod 777 /home/linaro/Desktop/golang/www/css/
chmod 777 /home/linaro/Desktop/golang/www/img/
chmod 777 /home/linaro/Desktop/golang/www/fonts/
chmod 777 /home/linaro/Desktop/golang/www/webfonts/
echo "set permission end"
