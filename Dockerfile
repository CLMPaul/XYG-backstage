FROM 172.28.82.183:30002/library/ubuntu:20.04 
MAINTAINER xuluo
ENV LOG_LEVEL=debug \
    DATABASE_SHOW_SQL=true \
    REDIS_ADDR="159.75.44.6:6379" \
    REDIS_PASSWORD="xueyi123gou" \
    MYSQL_ADDR="120.77.85.52:3306" \
    MYSQL_USER="xueyigou" \
    MYSQL_PASSWORD="Xueyigou123."
WORKDIR /tmp/
ADD ./make/main-linux  ./
RUN mkdir config && \
    chmod +x main-linux && \
    apt-get -qq update && \
    apt-get -qq install -y --no-install-recommends ca-certificates curl
COPY config /tmp/config
ENTRYPOINT ["/tmp/main-linux"]
