FROM python:3.11.0rc2-alpine3.16 as build

COPY . /code
WORKDIR /code
RUN apk update && apk add py3-pip
RUN pip3 install -r requirements.txt

CMD ["python3", "main.py"]
