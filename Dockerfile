FROM golang:1.10.0-alpine3.7 AS build

# Install tools required to build the project
# We need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/github.com/0xTanvir/pp/
WORKDIR /go/src/github.com/0xTanvir/pp/
# Install library dependencies
#RUN dep init
RUN dep ensure -vendor-only

# Copy all project and build it
# This layer is rebuilt when ever a file has changed in the project directory
COPY . /go/src/github.com/0xTanvir/pp/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pp .

# This results in a single layer image
FROM scratch
ADD etc /etc
COPY --from=build /go/src/github.com/0xTanvir/pp/pp /
CMD ["/pp", "run"]
EXPOSE 3030