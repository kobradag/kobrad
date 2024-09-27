ARG KODAPAD_IMAGE
ARG KODAPADMINER_IMAGE

FROM ${KODAPAD_IMAGE} as kobrad
FROM ${KODAPADMINER_IMAGE} as kobraminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/kobradag/kobrad

WORKDIR /go/src/github.com/kobradag/kobrad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=kobrad /app/ /app/
COPY --from=kobraminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/kobradag/kobrad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
