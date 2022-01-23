FROM python:3.9-slim-bullseye

COPY requirements.txt /requirements.txt
COPY app /app
COPY main.py /main.py

RUN pip install -r requirements.txt

EXPOSE 8080

CMD ["python3", "main.py"]