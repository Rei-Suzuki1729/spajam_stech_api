FROM postgres:latest

ENV LANG ja_JP.UTF-8
ENV LC_ALL "en_US.UTF-8"
ENV LC_TYPE "en_US.UTF-8"

#初期データ用のSQO追加
COPY database/initdb.d/init.sql /docker-entrypoint-initdb.d/
RUN chmod 755 /docker-entrypoint-initdb.d/init.sql
