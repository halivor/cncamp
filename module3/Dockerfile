FROM golang:1.19.2 AS build

ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /app
COPY module2/ /app/
RUN make

FROM centos:8.2.2004
COPY --from=build /app/bin/exec /usr/bin
ENTRYPOINT ["exec", "server"]
