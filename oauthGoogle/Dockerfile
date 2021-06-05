# images 
FROM golang

# craete app directory
RUN mkdir /app

# add all file to app directory
ADD . /app

# change directory to /app
WORKDIR /app

# create go build to compile app for production
RUN go build -o main .

# running command
CMD ["/app/main"]