FROM golang:1.11-stretch AS build

RUN go build -o /bin/notfound github.com/ken39arg/notfound

FROM stretch

COPY --from=build /bin/notfound ./notfound
ENV PORT=8080
ENTRYPOINT ["./notfound"]
