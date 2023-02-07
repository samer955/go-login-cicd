FROM golang:1.18-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o loginApp .

RUN chmod +x /app/loginApp

# build a tiny docker image

CMD [ "/app/loginApp" ]
