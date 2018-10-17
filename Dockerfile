FROM golang:1.11.1-alpine

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /go/src/project/
# Install library dependencies
RUN dep ensure -vendor-only

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY cmd/ Dockerfile Pkg/ /go/src/project/
RUN go build -o /bin/project cmd/main.go

## This results in a single layer image
#FROM scratch
#COPY --from=build /bin/project /bin/project
#ENTRYPOINT ["/bin/project"]
#CMD ["--help"]