# build stage
FROM golang:alpine as builder
RUN apk add git
RUN apk add make
WORKDIR /github.com/IamStubborN/calendar
COPY . .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN make all

# final stage
FROM alpine:latest
WORKDIR /root/petstore/
COPY --from=builder /github.com/IamStubborN/calendar .
COPY --from=builder /github.com/IamStubborN/calendar/config.yaml .

ENTRYPOINT ["./calendar"]
