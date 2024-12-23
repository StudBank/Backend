FROM golang:1.23.4-alpine AS build
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static -s -w" -o /app .

FROM scratch AS final
ARG env=dev

COPY --from=build /app /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY ./envs/$env .
ENTRYPOINT ["/app"]