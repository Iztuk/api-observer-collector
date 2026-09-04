FROM golang:1.26

# Destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the souce code
COPY . ./

# Build
RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build -o /api-observer-collector

# Run
CMD ["/api-observer-collector"]
