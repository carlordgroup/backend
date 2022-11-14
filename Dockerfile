FROM golang:1.19.3-buster

# set workdir
WORKDIR /app
# add the library first because the library won't usually change
ADD go.* ./
# download the go libs
RUN go mod download
# add other codes
ADD . /app
# build
RUN go build -o /app/main -buildvcs=false
RUN chmod +x /app/main

CMD ["/app/main"]