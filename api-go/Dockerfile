FROM golang
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .
EXPOSE 1323
CMD [ "/app/main" ]