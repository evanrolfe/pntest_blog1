FROM python:3.11-slim

ENV HOST_NAME=0.0.0.0
ENV PORT=8090

WORKDIR /app
ADD ./ /app

RUN apt-get update
RUN apt-get upgrade -qy
RUN apt-get install curl -qy

RUN pip install --upgrade pip
RUN pip install requests

EXPOSE 8090

ENTRYPOINT ["python", "/app/server.py"]
