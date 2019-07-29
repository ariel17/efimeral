FROM ubuntu:19.04
LABEL maintainer="ariel.gerardo.rios@gmail.com"

RUN ["apt", "update"]
RUN ["apt", "install", "-y", "linuxlogo"]
RUN ["apt", "install", "-y", "vim", "htop", "curl", "telnet"]
RUN ["apt", "install", "-y", "git", "python3.6", "golang-go", "openjdk-8-jdk", "nodejs", "npm"]

WORKDIR /tmp
RUN ["git", "clone", "https://github.com/liftoff/GateOne.git", ""]

WORKDIR /tmp/GateOne
RUN ["python", "setup.py", "install"]

ARG USER=efimeral
ARG PASSWD=efimeral
RUN useradd -m -p $(openssl passwd -1 ${PASSWD}) ${USER}
RUN echo linuxlogo >> /home/efimeral/.profile

EXPOSE 4200
# CMD ["shellinaboxd", "--disable-ssl", "--disable-ssl-menu", "--css=/etc/shellinabox/options-available/00_White On Black.css"]
