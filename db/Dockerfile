FROM mysql:8.0

ARG UID
ARG GID
ARG UNAME

ENV UID ${UID}
ENV GID ${GID}
ENV UNAME ${UNAME}

RUN groupadd -g ${GID} ${UNAME}
RUN useradd -u ${UID} -g ${UNAME} -m ${UNAME}

COPY ./script/container-init.sh /home/${UNAME}
RUN chmod +x /home/${UNAME}/container-init.sh

ENTRYPOINT ["sh", "-c", "/home/${UNAME}/container-init.sh"]
