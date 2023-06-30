FROM golang:alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY ./ ./

# Соберём приложение 
RUN go build ./cmd/umbrella
EXPOSE 8090
# Запустим приложение
CMD ["./umbrella"]