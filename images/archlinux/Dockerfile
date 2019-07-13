FROM archlinux/base:latest
LABEL maintainer="ariel.gerardo.rios@gmail.com"

RUN ["pacman", "-Syu", "--noconfirm"]
RUN ["pacman", "-S", "git", "sudo", "inetutils", "iputils", "base-devel", "--noconfirm"]
RUN ["pacman", "-Sy", "vim", "htop", "curl", "--noconfirm"]

COPY ["./build", "/tmp/build"]
COPY ["./install", "/tmp/install"]

RUN ["sudo", "-u", "nobody", "/tmp/build/linuxlogo.sh"]
RUN ["/tmp/install/linuxlogo.sh"]

RUN ["sudo", "-u", "nobody", "/tmp/build/shellinabox.sh"]
RUN ["/tmp/install/shellinabox.sh"]

# RUN ["pacman", "-Sy", "shellinabox", "openssl", "ca-certificates"]
# RUN ["pacman", "-Sy", "git", "python3.6", "golang-go", "openjdk-8-jdk", "nodejs", "npm"]
# 
# RUN echo linuxlogo >> /home/efimeral/.profile

ARG USER=efimeral
ARG PASSWD=efimeral
RUN useradd -s /bin/bash -m -p $(openssl passwd -1 ${PASSWD}) ${USER}

EXPOSE 4200
CMD ["shellinaboxd", "-u", "root", "-g", "root", "--no-beep", "-v", "-t", "-s", "/:SSH", "--css=/tmp/shellinabox-git/src/shellinabox/shellinabox/white-on-black.css"]
