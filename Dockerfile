FROM golang:1.18 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN useradd -u 10001 benthos

WORKDIR /workspace/benthos
# Update dependencies: On unchanged dependencies, cached layer will be reused
COPY go.* /workspace/benthos
RUN go mod download

# Build
COPY . /workspace/benthos
# Tag timetzdata required for busybox base image:
# https://github.com/benthosdev/benthos/issues/897
RUN go build -tags timetzdata -o benthos .

# Pack
FROM busybox AS package

LABEL maintainer="Mihai Todor <todormihai@gmail.com>"

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /workspace/benthos .

USER benthos

EXPOSE 4195

ENTRYPOINT ["/benthos"]
