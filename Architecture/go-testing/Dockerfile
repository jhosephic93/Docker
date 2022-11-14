FROM golang:1.13 as builder
WORKDIR /app
COPY ./ /app/
RUN go build main.go
#CMD ["/app/main"]

FROM scratch
COPY --from=builder /app/main /main
CMD ["/main"]
