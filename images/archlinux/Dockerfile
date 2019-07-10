FROM archlinux/base:latest
LABEL maintainer="ariel.gerardo.rios@gmail.com"

RUN ["pacman", "-Syu", "--noconfirm"]
RUN ["pacman", "-S", "git", "sudo", "inetutils", "--noconfirm"]

ARG USER=efimeral
ARG PASSWD=efimeral
RUN useradd -m -p $(openssl passwd -1 ${PASSWD}) ${USER}

COPY ["./install", "/tmp/install"]
RUN ["chmod", "+x", "/tmp/install/linuxlogo.sh"]
RUN ["sudo", "-u", "efimeral", "/tmp/install/linuxlogo.sh"]

RUN ["pacman", "-Sy", "vim", "htop", "curl", "telnet"]
# RUN ["pacman", "-Sy", "shellinabox", "openssl", "ca-certificates"]
# RUN ["pacman", "-Sy", "git", "python3.6", "golang-go", "openjdk-8-jdk", "nodejs", "npm"]
# 
# RUN echo linuxlogo >> /home/efimeral/.profile
# 
# EXPOSE 4200
# CMD ["shellinaboxd", "--disable-ssl", "--disable-ssl-menu", "--css=/etc/shellinabox/options-available/00_White On Black.css"]
