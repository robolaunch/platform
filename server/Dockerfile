FROM golang:1.21 AS build
WORKDIR /go/src
COPY . .
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o server .

FROM gcr.io/distroless/static:nonroot AS runtime
COPY --from=build /go/src/server ./
EXPOSE 8080/tcp
ENTRYPOINT ["./server"]
