FROM python:3.11-slim

WORKDIR /usr/src/app

COPY . .

RUN pip install flask

EXPOSE 8030

CMD ["python", "app.py"]