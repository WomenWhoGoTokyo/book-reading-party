FROM golang:1.18-alpine AS build
WORKDIR /go/src/proglog
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/proglog ./cmd/proglog

FROM scratch
COPY --from=build /bin/proglog /bin/proglog
ENTRYPOINT ["/bin/proglog"]
