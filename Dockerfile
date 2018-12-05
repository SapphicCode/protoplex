# build
FROM alpine:edge AS build

RUN apk update && apk add musl-dev go
COPY protoplex/ /tmp/protoplex/protoplex
COPY protoplex.go /tmp/protoplex/
RUN go build /tmp/protoplex/protoplex.go

# deploy
FROM alpine:latest
COPY --from=build /protoplex /protoplex

USER 999
ENTRYPOINT ["/protoplex"]
EXPOSE 8443/tcp
STOPSIGNAL SIGINT
