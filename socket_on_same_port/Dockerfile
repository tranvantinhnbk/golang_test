FROM golang:1.24-alpine

# Install Python3, pip, and iproute2
RUN apk add --no-cache iproute2 iputils

# Copy application scripts
WORKDIR /app
COPY server.go client.go ./

# Make scripts executable
RUN chmod +x server.go client.go

# Default command (overridden per service)
CMD ["/bin/sh"]