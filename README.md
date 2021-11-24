# Kong and Consul Lab

## Prerequisites

1. docker

## Deploy Consul Server
```
docker run -d -p 8500:8500 -p 8600:8600/udp --name=consul-server consul agent -server -ui -node=consul-server-1 -bootstrap-expect=1 -client=0.0.0.0  -log-level=debug
```


## Deploy Home APP
```
docker build -t home-image .
docker run -d -p 8070:8070 --name=home-app home-image
```

## Deploy Product APP
```
docker build -t product-image .
docker run -d -p 8090:8090 --name=product-app product-image
```

## Deploy Career APP
```
docker build -t career-image .
docker run -d -p 8080:8080 --name=career-app career-image
```

## Test DNS
please refer to this doc: https://www.consul.io/docs/discovery/dns
```
dig @127.0.0.1 -p 8600 home-app.service.consul ANY
dig @127.0.0.1 -p 8600 career-app.service.consul ANY
dig @127.0.0.1 -p 8600 product-app.service.consul ANY
```


## Deploy Kong

1. deploy postgress
```
docker run -d --name kong-database \
  -p 5432:5432 \
  -e "POSTGRES_USER=kong" \
  -e "POSTGRES_DB=kong" \
  -e "POSTGRES_PASSWORD=kong" \
  postgres:9.6
```

2. Migrate postgress
```
docker run --rm --link kong-database:kong-database \
  -e "KONG_DATABASE=postgres" \
  -e "KONG_PG_HOST=kong-database" \
  -e "KONG_PG_PASSWORD=kong" \
  -e "KONG_PASSWORD=kong" \
  kong:latest kong migrations bootstrap
```

3. Start Kong gateway
```
docker run -d --name kong \
  --link kong-database:kong-database\
  -e "KONG_DATABASE=postgres" \
  -e "KONG_PG_HOST=kong-database" \
  -e "KONG_PG_USER=kong" \
  -e "KONG_PG_PASSWORD=kong" \
  -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
  -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
  -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
  -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
  -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
  -e "KONG_DNS_RESOLVER=172.17.0.2:8600" \
  -p 8000:8000 \
  -p 8443:8443 \
  -p 8001:8001 \
  -p 8444:8444 \
  kong:latest
```

## Add service and routes to KONG

Please refer to this doc: https://docs.konghq.com/gateway-oss/2.5.x/admin-api/#service-object
```
curl -i -X POST \
--url http://localhost:8001/services/ \
--data 'name=home-svc' \
--data 'host=home-app.service.consul'

curl -i -X POST \
--url http://localhost:8001/services/ \
--data 'name=product-svc' \
--data 'host=product-app.service.consul' \
--data 'path=/product'

curl -i -X POST \
--url http://localhost:8001/services/ \
--data 'name=career-svc' \
--data 'host=career-app.service.consul' \
--data 'path=/career'

curl -i -X POST \
--url http://localhost:8001/services/ \
--data 'name=blog-svc' \
--data 'host=career-app.service.consul' \
--data 'path=/blog'
```

Please refer to this doc: https://docs.konghq.com/gateway-oss/2.5.x/admin-api/#route-object
```
curl -i -X POST \
--url http://localhost:8001/services/home-svc/routes \
--data 'hosts[]=localhost' \
--data 'paths=/home'

curl -i -X POST \
--url http://localhost:8001/services/product-svc/routes \
--data 'hosts[]=localhost' \
--data 'paths=/product'

curl -i -X POST \
--url http://localhost:8001/services/career-svc/routes \
--data 'hosts[]=localhost' \
--data 'paths=/career'

curl -i -X POST \
--url http://localhost:8001/services/blog-svc/routes \
--data 'hosts[]=localhost' \
--data 'paths=/blog'
```


## Deploy Konga

```
docker run --rm --link kong-database:kong-database \
pantsel/konga -c prepare -a postgres -u postgresql://kong:kong@kong-database:5432/konga_db

docker run -d --name konga \
-p 1337:1337 \
--link kong-database:kong-database \
-e "DB_ADAPTER=postgres" \
-e "DB_HOST=kong-database" \
-e "DB_USER=kong" \
-e "DB_PASSWORD=kong" \
-e "DB_DATABASE=konga_db" \
-e "KONGA_HOOK_TIMEOUT=120000" \
-e "NODE_ENV=production" \
pantsel/konga
```