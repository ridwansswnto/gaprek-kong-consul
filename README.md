# Kong and Consul Lab

## Prerequisites

1. docker

## Deploy Consul Server
```
docker run -d -p 8500:8500 -p 8600:8600/udp --name=consul-server consul agent -server -ui -node=consul-server-1 -bootstrap-expect=1 -client=0.0.0.0 
```


## Deploy Home APP
```
docker build -t home-image .
docker run -d -p 8070:8070 --name=home-app home-image
```


### Deploy Kong

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
  -e "KONG_DNS_RESOLVER=172.17.0.3:8600" \
  -p 8000:8000 \
  -p 8443:8443 \
  -p 8001:8001 \
  -p 8444:8444 \
  kong:latest
```

## Add service to KONG
```
curl -i -X POST \
--url http://localhost:8001/services/ \
--data 'name=home-svc' \
--data 'host=home-app.service.consul'
```

```
curl -i -X POST \
--url http://localhost:8001/services/home-svc/routes \
--data 'hosts[]=localhost' \
--data 'paths=/home'
```