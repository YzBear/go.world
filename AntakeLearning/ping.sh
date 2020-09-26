#!/bin/bash
while getopts ":w:" opt
do
    case $opt in
        w)
        ping -c 3 $OPTARG | awk '/time=/{print $8}' | awk -F "=" '{sum+=$2} END{printf("平均延迟为：%f\n",sum/3)}' 
        ;;
        ?)
        echo "未知参数"
        exit 1;;
    esac
done
