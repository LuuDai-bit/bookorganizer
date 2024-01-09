FROM golang:1.21-bookworm

ARG APP_PATH=/book-organizer

RUN apt-get update -qq && apt-get install sudo

RUN mkdir ${APP_PATH}
WORKDIR ${APP_PATH}

COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]
