# geometry-server
#
# VERSION 1.0

FROM ubuntu

MAINTAINER Matthew McFarland mmcfarland@gmail.com

# Install up to date dependencies
RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y libgeos-dev wget git-core build-essentials

# Install and setup the go language tools
RUN wget https://go.googlecode.com/files/go1.1.2.linux-amd64.tar.gz --no-check-certificate
RUN tar -C /usr/local -xzf go1.1.2.linux-amd64.tar.gz 
RUN mkdir $HOME/go
ENV GOPATH ${HOME}go
RUN echo $GOPATH

# Get golang bindings for GEOS
RUN /usr/local/go/bin/go get github.com/paulsmith/gogeos/geos
