FROM tf-go-base:1.0
WORKDIR /go/src/imgrecognition
COPY . .
RUN go build
RUN go test -v
ENTRYPOINT [ "/go/src/imgrecognition/imgrecognition" ]
