FROM golang as builder

ENV GO111MODULE=on

WORKDIR soron

COPY . .

RUN go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

#-ldflags "-s -w"
RUN go build -o soron-sp

FROM alpine

WORKDIR /app

COPY --from=builder /go/soron/soron-sp .

ENTRYPOINT ["./soron-sp"]