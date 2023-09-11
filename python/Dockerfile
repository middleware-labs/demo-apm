# docker build -t ghcr.io/middleware-labs/demo-apm-python .
# docker push ghcr.io/middleware-labs/demo-apm-python
# export PATH=$PATH:/home/acq053/.local/bin

# Use an official Python runtime as the base image
FROM python:3.9

# Set the working directory in the container
WORKDIR /app

# Copy the contents of the local project directory into the container
COPY . .

# Create a virtual environment
RUN python -m venv venv

# Activate the virtual environment
RUN export PATH="$PATH:/app/venv/bin"

# RUN /app/venv/bin/python -m pip install --upgrade pip

# Install the project dependencies using pip
RUN pip install --no-cache-dir -r requirements.txt
RUN middleware-bootstrap -a install
RUN apt-get update; apt-get install -y curl
# # # Specify the command to run the application
ENTRYPOINT [ "sh", "/app/entrypoint.sh"]

#RUN pip3 install -r requirements.txt
# RUN pip install setuptools wheel twine
