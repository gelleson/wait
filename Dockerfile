# Dockerfile
FROM alpine:3.14.3
COPY wait /usr/bin/wait
ENTRYPOINT ["/usr/bin/wait"]