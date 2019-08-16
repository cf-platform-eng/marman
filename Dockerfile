FROM ubuntu:xenial
LABEL maintainer="Pivotal Platform Engineering ISV-CI Team <cf-isv-dashboard@pivotal.io>"

COPY build/marman-linux /usr/local/bin/marman

ENTRYPOINT [ "marman" ]