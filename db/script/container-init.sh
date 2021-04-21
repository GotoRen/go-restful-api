#!/bin/sh

echo ${GID}
echo ${UID}
echo ${UNAME}

if [ ! -e '/check' ]; then
    touch /check

    echo "Runninng Setup."
    groupadd -g ${GID} ${UNAME}
    useradd -u ${UID} -g ${UNAME} -m ${UNAME}
    chown ${UNAME}:${UNAME} -R /var/lib/mysql
    echo "Complete Owner Setup!!"
else
    echo "Already Setup." 
    chown ${UNAME}:${UNAME} -R /var/lib/mysql
    echo "Complete Owner Setup!!"
fi
