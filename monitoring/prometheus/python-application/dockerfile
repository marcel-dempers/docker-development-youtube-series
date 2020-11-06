FROM python:3.7.3-alpine3.9 as prod

RUN mkdir /app/
WORKDIR /app/

COPY ./src/requirements.txt /app/requirements.txt
RUN pip install -r requirements.txt

COPY ./src/ /app/

ENV FLASK_APP=server.py
CMD flask run -h 0.0.0 -p 5000
