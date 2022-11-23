FROM golang

WORKDIR /app
COPY . .

RUN go build -o app . 

CMD [ "./app" ]

EXPOSE 1323