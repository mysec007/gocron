# 内部镜像构建部署
FROM docker.17usoft.com/anquan/centos-gocron:v1.1
#FROM docker.17usoft.com/base/golang:1.15

RUN go build .


COPY releases/gocron-linux-amd64/gocron /app/gocron
COPY releases/gocron-node-linux-amd64/gocron_node_linux /app/gocron_node_linux
COPY releases/start.sh /app/start.sh


WORKDIR /app
RUN mkdir conf \
 && chmod 777 gocron \
 && chmod 777 gocron_node_linux \
 && chmod 777 start.sh \

COPY releases/app_test.ini /app/conf/app.ini
COPY releases/install.lock /app/conf/install.lock


ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

EXPOSE 80
EXPOSE 5921
CMD sh -c "/app/start.sh"
#/app/gocron web --host 0.0.0.0 --port 80 &
#su root -c '/app/gocron_node_linux -s 0.0.0.0:5921'

# 处理agent端依赖
#连接git
COPY releases/conf/.ssh /root/.ssh/
RUN chmod 600 /root/.ssh/id_rsa && chmod 644 /root/.ssh/id_rsa.pub
#python依赖
RUN pip3 install --upgrade pip && pip3 install -r /opt/iast/webapi/requirements.txt