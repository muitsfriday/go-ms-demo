FROM golang:1.11.2

RUN go get -u github.com/gin-gonic/gin
RUN go get gopkg.in/mgo.v2
RUN go get github.com/oxequa/realize
RUN go get github.com/dgrijalva/jwt-go

ADD ./ /src/app

WORKDIR /src/app


CMD ["go", "run", "main.go"]