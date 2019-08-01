# FROM golang:1.12-alpine AS builder
# ARG basedir=/go/build
# WORKDIR ${basedir}

# # Copy src and build
# ADD . ${basedir}
# RUN go build -o app

FROM alpine:3.8
# for https
# RUN apk add --no-cache ca-certificates
# Finally we copy the statically compiled Go binary.
# COPY --from=builder /go/build/app /app
COPY target/test /opt/
ENTRYPOINT ["/opt/test"]
