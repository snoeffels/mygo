FROM golang:1.21

WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify

copy .env-docker .env
CMD ["go", "run", "." ]