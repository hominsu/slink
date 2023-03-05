# generate stage
FROM golang:alpine AS generator

# install deps
RUN apk --no-cache add bash build-base curl findutils git protobuf protobuf-dev tar

COPY .. /src
WORKDIR /src

# install buf
RUN PREFIX="/usr/local" && \
    VERSION="1.13.1" && \
    curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m).tar.gz" | \
    tar -xvzf - -C "${PREFIX}" --strip-components 1

# workspace
RUN if [ ! -f "go.work" ]; then \
    go work init && go work use -r ./app && go mod tidy; \
    fi

# generate all
RUN make init && make api && make conf && make ent && make wire

# build stage
FROM golang:latest AS builder
# app dir
ARG APP_RELATIVE_PATH

# build app
COPY --from=generator /src /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

# build app
RUN make dep && make build

# final stage
FROM debian:stable-slim
# image info
ARG AUTHOR_NAME
ARG AUTHOR_EMAIL
ARG VERSION
# label
LABEL author=${AUTHOR_NAME} email=${AUTHOR_EMAIL} version=${VERSION}
# app dir
ARG APP_RELATIVE_PATH

# install: ca-certificates netbase
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

# cpoy the app from builder
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
WORKDIR /app
EXPOSE 8000
VOLUME /data/conf
CMD ["./server", "-conf", "/data/conf"]
