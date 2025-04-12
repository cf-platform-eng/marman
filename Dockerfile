ARG ubuntu_image=tas-ecosystem-docker-virtual.usw1.packages.broadcom.com/ubuntu

FROM ${ubuntu_image}
LABEL maintainer="Pivotal Platform Engineering ISV-CI Team <cf-isv-dashboard@pivotal.io>"

RUN apt-get update && apt-get -y install ca-certificates && rm -rf /var/lib/apt/lists/*

COPY build/marman-linux /usr/local/bin/marman

ENTRYPOINT [ "marman" ]