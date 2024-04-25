FROM python:latest

RUN pip install jsonsubschema
RUN pip install pyyaml
COPY entrypoint.py /entrypoint.py

ENTRYPOINT ["python", "/entrypoint.py"]