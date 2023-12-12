FROM golang:1.21-bookworm

ARG USERID
ARG GROUPID
ARG APP_PATH=/book-organizer

RUN apt-get update -qq && apt-get install sudo

RUN mkdir ${APP_PATH}
WORKDIR ${APP_PATH}

RUN useradd -ms /bin/bash bo --no-log-init
RUN usermod -u ${USERID} bo
RUN groupmod -g ${GROUPID} bo || exit 0

RUN echo "bo ALL=(ALL:ALL) NOPASSWD: ALL" > /etc/sudoers.d/bo
RUN chown -R bo:bo ${APP_PATH}

USER bo
COPY --chown=bo:bo . .

COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]
