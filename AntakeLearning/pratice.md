> linux 提取本机ip

```shell

[root@iZwz9almo8p830btq7voo9Z ~]# ifconfig | awk '/broadcast/{print $0}' | awk '{print $2}' | awk -F "." '{if($3!=0 && $4!=1){print $0}}'
172.18.xxx.xxx
```

