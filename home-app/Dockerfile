FROM consul
RUN apk add go supervisor
RUN mkdir /build
ADD home-app.go go.mod /build/
ADD home-app.json /consul/config/home-app.json
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf
WORKDIR /build
RUN go build
ENTRYPOINT ["supervisord","-c","/etc/supervisor/conf.d/supervisord.conf"]