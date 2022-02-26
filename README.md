# homework-1-ErdemOzgen
![PicusPatika](./img/picuspatika.png)

This is first assignment for picus patika.dev.

# How to install ?

```bash
go get ./...
go install 
homework-1-ErdemOzgen

```

# How it works ?

![img1](./img/1.png)







### Docker file for docker buildings 

```bash

docker build -t homework-1-ErdemOzgen .

```


```dockerfile
FROM golang

WORKDIR $GOPATH/src/homework-1-ErdemOzgen

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o .

EXPOSE 8000:8000

CMD ["homework-1-ErdemOzgen server"]

```
