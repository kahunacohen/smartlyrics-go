FROM golang:1.17-buster AS build
WORKDIR /app

COPY . ./
RUN go mod download
RUN go build -o /smartlyrics .

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /smartlyrics /smartlyrics
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/smartlyrics"]
