FROM golang AS build
WORKDIR /go/src/github.com/postmates/butte
COPY butte.go .
RUN go get -d .
RUN go build -o butte .

FROM gcr.io/distroless/base
COPY --from=build /go/src/github.com/postmates/butte/butte .
EXPOSE 2112
CMD ["./butte"]
