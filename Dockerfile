FROM golang:1.17.3-buster as builder

WORKDIR /app

COPY . .

RUN go install

ENV GOOS linux
ENV GOARCH arm64
ENV CGO_ENABLED=0
ENV EXTRA_LDFLAGS=""

RUN CGO_ENABLED=$CGO_ENABLED GOOS=linux GOARCH=$GOARCH go build $EXTRA_LDFLAGS -o wait main.go

FROM alpine:3.14.3 as worker
COPY --from=builder /app/wait /usr/bin/wait

RUN chmod +x /usr/bin/wait && ls

CMD ["/usr/bin/wait"]