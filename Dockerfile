FROM golang:1.11.1-alpine

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git
RUN alias go='http_proxy=http://ctu-net-mwg:3128 https_proxy=http://ctu-net-mwg:312 no_proxy=localhost,127.0.0.0/8,::1 go'
RUN go get github.com/golang/dep/cmd/dep

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/github.com/steven-liao/objectserver/
WORKDIR /go/src/github.com/steven-liao/objectserver/
# Install library dependencies
RUN dep ensure -vendor-only

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY Dockerfile . 

RUN mkdir cmd pkg
COPY cmd cmd/
COPY pkg pkg/

RUN go build -o /bin/project cmd/main.go
ENTRYPOINT ["/bin/project"]
EXPOSE 8080
## This results in a single layer image
#FROM scratch
#COPY --from=build /bin/project /bin/project
#ENTRYPOINT ["/bin/project"]
#CMD ["--help"]
