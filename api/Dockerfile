FROM docker:latest
LABEL maintainer="Ariel Gerardo Ríos <arielgerardorios@gmail.com>"

WORKDIR /app
COPY ["./api", "."]

EXPOSE 8080

ENV ENVIRONMENT production
ENV BASE_URL http://localhost

CMD ["./api"]
