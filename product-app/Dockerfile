FROM consul
RUN apk add go supervisor
RUN mkdir /build
ADD product-app.go go.mod /build/
ADD product-app.json /consul/config/product-app.json
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf
WORKDIR /build
RUN go build
ENTRYPOINT ["supervisord","-c","/etc/supervisor/conf.d/supervisord.conf"]