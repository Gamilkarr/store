FROM golang:latest AS build-stage
WORKDIR /app
COPY go.mod ./
RUN go env
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build ./cmd/store/main.go

FROM build-stage
RUN go test -race -v ./...

FROM scratch
WORKDIR /
COPY --from=build-stage /app/main /store
ENTRYPOINT ["/store"]