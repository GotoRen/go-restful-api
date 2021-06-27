#!/bin/sh

ENV_VOL=".env.vol"

if [ -e $ENV_VOL ]; then
    rm $ENV_VOL
    echo "[INFO] Deleted .env.vol"
fi

echo "LINUX_MYSQL_UID=$(id -u $USER)" >> $ENV_VOL
echo "LINUX_MYSQL_GID=$(id -g $USER)" >> $ENV_VOL

echo "[INFO] Created .env.vol"