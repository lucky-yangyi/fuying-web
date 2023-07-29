FROM registry.cn-hangzhou.aliyuncs.com/fxy-zhike/debian:buster

RUN mkdir -p /cmd
WORKDIR  /cmd
COPY fuying-web  .
COPY conf/app.json ./conf/app.json
EXPOSE 8888
CMD ["./fuying-web"]
