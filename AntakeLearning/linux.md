```
# 查看系统磁盘占用
df -h | grep "/$" | awk '{print $(NF-1)}' | awk -F "%" '{print $1}'
```
