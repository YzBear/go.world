#!/bin/bash
Db=
User=
Pass=
Url=
while getopts ":d:u:p:h:U:" opt
do
    case $opt in
        d)
        Db=$OPTARG
        ;;
		u)
        User=$OPTARG
        ;;
		U)
        Url=$OPTARG
        ;;
		p)
        Pass=$OPTARG
        ;;
		h|?)
        echo "-h help"
		echo "-d dbname"
        echo "-u username"
        echo "-p password"
		echo "-U url"
		exit 1
        ;;
    esac
done
if [ ! $Db ] || [ ! $User ] || [ ! $Pass ] || [ ! $Url ]
then
	echo "参数不完整"
	echo "-h help"
	exit 1
fi
Mysql_Dump=$(which mysqldump)
Backup_Dir=/backup
Date=$(date +%F)
Backup_Path=$Backup_Dir/$Date
test -d $Backup_Path || mkdir -p $Backup_Path
if [ $? -eq 0 ]
then 
	echo "=============Backup Is Starting=============="
	$Mysql_Dump -u$User -h $Url -p$Pass -B $Db > $Backup_Path/mysql_db.sql
	if [ $? -eq 0 ]
	then
		echo "==============Backup Is Done==============="
	else
		echo "============Beck Error==============="
	fi
fi
