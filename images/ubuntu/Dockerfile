FROM ubuntu:19.04
LABEL maintainer="ariel.gerardo.rios@gmail.com"

RUN ["apt", "update"]
RUN ["apt", "install", "-y", "linuxlogo"]
RUN ["apt", "install", "-y", "vim", "htop", "curl", "telnet"]
RUN ["apt", "install", "-y", "git", "golang-go", "openjdk-8-jdk", "nodejs", "npm", "gcc"]
RUN ["apt", "install", "-y", "python-pip", "python-dev", "python3", "python3-pip", "python3-dev"]

WORKDIR /tmp
ADD ./gateone.sh .
RUN ["./gateone.sh"]

ARG USER=efimeral
ARG PASSWD=efimeral
RUN useradd -m -p $(openssl passwd -1 ${PASSWD}) ${USER}
RUN echo linuxlogo >> /home/efimeral/.profile

WORKDIR /
EXPOSE 10443
CMD ["gateone", "--disable-ssl", "--port=10443"]
