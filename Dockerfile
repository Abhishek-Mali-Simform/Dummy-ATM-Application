FROM golang:1.20

WORKDIR /app

COPY . .

RUN go get github.com/beego/bee/v2@latest
RUN go install github.com/beego/bee/v2@latest

RUN GO111MODULE=on

RUN go get
RUN go build -o /Application

EXPOSE 5000

CMD ["bee","run","-gendoc=true","-downdoc=true"]