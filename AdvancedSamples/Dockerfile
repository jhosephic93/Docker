# Use a specific version of Python base image for better consistency
FROM python:3.9-slim
# Set a working directory to avoid using the root directory
WORKDIR /app
# Copy only the requirements file to leverage Docker cache
COPY requirements.txt ./
# Install dependencies in a way that makes the image more secure and lightweight
RUN pip install --no-cache-dir -r requirements.txt
# Copy only the necessary source code file
COPY webpage.py ./webpage.py
# Expose the port on which the application will run
EXPOSE 8080
# Use array syntax for CMD and specify the full path for clarity
CMD ["python", "/app/webpage.py"]
# Create a non-root user and switch to it for better security
RUN useradd -m myuser
USER myuser