FROM consul
RUN apk add go supervisor
RUN mkdir /build
ADD career-app.go go.mod /build/
ADD career-app.json /consul/config/career-app.json
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf
WORKDIR /build
RUN go build
ENTRYPOINT ["supervisord","-c","/etc/supervisor/conf.d/supervisord.conf"]