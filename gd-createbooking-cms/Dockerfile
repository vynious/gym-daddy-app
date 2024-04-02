FROM python:3-slim
WORKDIR /usr/src/app
COPY requirements.txt ./
RUN python -m pip install --no-cache-dir -r requirements.txt

# Copy the entire 'pb' directory into the image
COPY pb/ pb/
# COPY .gitignore .
EXPOSE 5002

COPY create_booking.py .
CMD [ "python", "./create_booking.py" ]
