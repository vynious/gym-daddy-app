FROM python:3-slim
WORKDIR /usr/src/app
COPY requirements.txt ./
RUN python -m pip install --no-cache-dir -r requirements.txt

# Copy the entire 'pb' directory into the image
COPY pb/ pb/
COPY invokes.py .
# COPY .env .
EXPOSE 5001

COPY create_class.py .
CMD [ "python", "./create_class.py" ]
