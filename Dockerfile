FROM golang:alpine as builder
MAINTAINER Jessica Frazelle <jess@linux.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	ca-certificates

COPY . /go/src/github.com/jessfraz/party-clippy

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		git \
		gcc \
		libc-dev \
		libgcc \
		make \
	&& cd /go/src/github.com/jessfraz/party-clippy \
	&& make static \
	&& mv party-clippy /usr/bin/party-clippy \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

FROM scratch

COPY --from=builder /usr/bin/party-clippy /usr/bin/party-clippy
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENTRYPOINT [ "party-clippy" ]
CMD [ "--help" ]
