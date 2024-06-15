FROM python:3.11.9-slim

WORKDIR /code

COPY requirements.txt requirements.txt
RUN --mount=type=cache,target=/root/.cache/pip pip3 install -r requirements.txt
COPY . /code