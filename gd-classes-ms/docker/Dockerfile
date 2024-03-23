FROM python:3-slim
WORKDIR /usr/src/app
COPY requirements.txt ./
RUN python -m pip install --no-cache-dir -r requirements.txt
COPY ./classes.py .
CMD [ "python", "./classes.py" ]


    # docker run -p 5200:5200 -e dbURL=mysql+mysqlconnector://is213@host.docker.internal:8889/classes camillechua/class:1.0