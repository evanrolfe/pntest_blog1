FROM python:3.11-slim

ENV PROXY_PORT=8091
ENV WEB_UI_PORT=8092

WORKDIR /app
ADD ./ /app

RUN apt-get update
RUN apt-get upgrade -qy
RUN apt-get install iptables curl -qy

RUN pip install --upgrade pip
RUN pip install mitmproxy

RUN useradd --create-home mitmproxyuser
ENTRYPOINT /bin/bash run.sh
