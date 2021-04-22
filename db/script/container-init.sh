#!/bin/sh

echo ${GID}
echo ${UID}
echo ${UNAME}

if [ ! -e '/check' ]; then
    touch /check

    echo "Runninng Setup."
    chown ${UNAME}:${UNAME} -R /var/lib/mysql
else
    echo "Already Setup." 
fi

/bin/bash