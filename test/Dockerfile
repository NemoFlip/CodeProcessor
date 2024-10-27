FROM python:3.9-slim

WORKDIR testPy

RUN pip install pytest requests

COPY test/tests.py .

CMD ["pytest", "tests.py"]