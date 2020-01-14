FROM golang:1.13.6

ARG KOALA_VERSION=1.0.6

LABEL maintainer="Clivern <hello@clivern.com>"

WORKDIR /app

RUN curl -sL https://github.com/Clivern/Koala/releases/download/${KOALA_VERSION}/Koala_${KOALA_VERSION}_Linux_x86_64.tar.gz | tar xz

RUN mv Koala koala

RUN rm LICENSE

RUN rm README.md

EXPOSE 8080

CMD ["./koala"]
