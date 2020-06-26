FROM ubuntu:20.04
LABEL maintainer="ariel.gerardo.rios@gmail.com"

ARG DEBIAN_FRONTEND=noninteractive

RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "linuxlogo"]
RUN ["apt-get", "install", "-y", "vim", "htop", "curl", "telnet"]
RUN ["apt-get", "install", "-y", "shellinabox", "openssl", "ca-certificates"]
RUN ["apt-get", "install", "-y", "git", "python3.6", "golang-go", "openjdk-8-jdk", "nodejs", "npm"]

ARG USER=efimeral
ARG PASSWD=efimeral
RUN useradd -m -p $(openssl passwd -1 ${PASSWD}) ${USER}
RUN echo linuxlogo >> /home/efimeral/.profile

EXPOSE 4200
CMD ["shellinaboxd", "--disable-ssl", "--disable-ssl-menu", "--css=/etc/shellinabox/options-available/00_White On Black.css"]
