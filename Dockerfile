FROM golang:1.19 as build

COPY . /app
WORKDIR /app
RUN go build .

FROM alpine:latest as app

RUN apk --no-cache add bash libc6-compat && \
    mkdir /app && \
    mkdir /app/web

COPY --from=build /app/portal /app
COPY --from=build /app/web/templates /app/web/templates
WORKDIR /app

CMD [ "/app/portal" ]