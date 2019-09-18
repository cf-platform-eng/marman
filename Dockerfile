FROM ubuntu
LABEL maintainer="Pivotal Platform Engineering ISV-CI Team <cf-isv-dashboard@pivotal.io>"

RUN apt-get update && apt-get -y install ca-certificates && rm -rf /var/lib/apt/lists/*

COPY build/marman-linux /usr/local/bin/marman

ENTRYPOINT [ "marman" ]