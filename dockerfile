FROM golang:1.23.0-alpine

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -v -o main .

CMD ["./main"]


# FROM golang:1.23-alpine

# WORKDIR /app
# # Create the directory and set permissions
# RUN mkdir -p /app/images && chmod -R 777 /app/images
# COPY go.mod go.sum ./

# RUN go mod download
# COPY . .
# RUN go build -v -o main ./cmd


# RUN adduser -D nonroot && mkdir -p /etc/sudoers.d \
#         && echo "nonroot ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/nonroot \
#         && chmod 0440 /etc/sudoers.d/nonroot

# USER nonroot

# CMD ["./main"]
