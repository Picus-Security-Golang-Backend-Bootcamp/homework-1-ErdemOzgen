# homework-1-ErdemOzgen
![PicusPatika](./img/picuspatika.png)

homework-1-ErdemOzgen created by GitHub Classroom


### Docker file for docker buildings 

```dockerfile
FROM golang

WORKDIR $GOPATH/src/src/homework-1-ErdemOzgen

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o .

EXPOSE 8000

CMD ["homework-1-ErdemOzgen"]

```
