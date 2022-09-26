FROM python:3.11.0rc2-alpine3.16 as build

WORKDIR /code
RUN apk update && apk add py3-pip
COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt
COPY . .
RUN apk add chromium

CMD ["python3", "main.py"]
