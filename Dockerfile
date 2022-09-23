FROM golang
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /restaurant ./cmd/app/main.go
EXPOSE 8080
CMD [ "/restaurant" ]