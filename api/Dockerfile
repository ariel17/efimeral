FROM golang:1.12.5
LABEL maintainer="Ariel Gerardo Ríos <arielgerardorios@gmail.com>"

WORKDIR /app
COPY ["./api", "."]

EXPOSE 8080

ENV ENVIRONMENT production

CMD ["./api"]
