#!/bin/bash
redis-server  /home/itheima/workspace/go/src/FullHouse/conf/redis.conf
fdfs_trackerd /home/itheima/workspace/go/src/FullHouse/conf/tracker.conf
fdfs_storaged /etc/fdfs/storage2.conf

