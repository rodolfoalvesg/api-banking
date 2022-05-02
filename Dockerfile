FROM golang:latest
RUN mkdir /app-banking
COPY . /app-banking
WORKDIR /app-banking
RUN go build -o /api main.go
EXPOSE 3000
CMD ["/api"]