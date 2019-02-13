FROM golang:1.12beta2-stretch

RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt-get update && apt-get install -y nodejs
RUN npm update && npm i -D electron@beta

ENV GOOS js
ENV GOARCH wasm
ENV PATH "${PATH}:/usr/local/go/misc/wasm"

WORKDIR /go/src/github.com/madlambda/Nine

RUN go get -v github.com/madlambda/spells/...

COPY . .