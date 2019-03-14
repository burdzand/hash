FROM python:3.6-alpine3.6

LABEL client.api="0.1.0"
ADD . /opt/client
RUN apk add --no-cache tzdata
ENV TZ America/Sao_Paulo

# alpine-sdk
RUN apk add --update bash
RUN apk add --no-cache --virtual .build-deps \
 build-base mariadb-dev libffi-dev gcc libc-dev linux-headers python3-dev \
   && pip install -r /opt/client/requirements.txt \
   && find /usr/local \
       \( -type d -a -name test -o -name tests \) \
       -o \( -type f -a -name '*.pyc' -o -name '*.pyo' \) \
       -exec rm -rf '{}' + \
   && runDeps="$( \
       scanelf --needed --nobanner --recursive /usr/local \
               | awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' \
               | sort -u \
               | xargs -r apk info --installed \
               | sort -u \
   )" \
   && apk add --virtual .rundeps $runDeps  \
   && apk del .build-deps mariadb-dev libffi-dev gcc libc-dev build-base python3-dev
RUN rm -Rf ~/.cache && rm -rf /var/cache/apk/* && rm -rf /tmp/*
EXPOSE 8000