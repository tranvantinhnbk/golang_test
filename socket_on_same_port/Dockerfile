FROM golang:1.24-alpine

# Install Python3, pip, and iproute2
RUN apk add --no-cache iproute2=6.1.0-r2 iputils=20221126-r2

# Copy application scripts
WORKDIR /app
COPY server.py client.py ./

# Make scripts executable
RUN chmod +x server.py client.py

# Default command (overridden per service)
CMD ["/bin/sh"]