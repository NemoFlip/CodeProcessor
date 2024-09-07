FROM python:latest

WORKDIR testPy

RUN pip install pytest
RUN pip install requests

COPY test/tests.py .

CMD ["pytest", "tests.py"]