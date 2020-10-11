## 1 Run database docker container

docker run --name ip2proxy -d -e TOKEN=<your_ip2location_api_key> -e CODE=PX7LITECSV -e MYSQL_PASSWORD=<your_password> ip2proxy/mysql

## 2 Set config file: ./config/dev_config.json and change mysql password

## 3 Build locally

Normal docker build (not using modules) - image size is 775MB

`docker build . -t go-ip2proxy-api-full`

Optimized Build (with modules - image size is 389MB

`docker build . -f Dockerfile.mod -t go-ip2proxy-api-modules`

Multi-stage build (fully optimized) - image size is 16 MB

`docker build . -f Dockerfile.multistage -t go-ip2proxy-api-multi`

## 4 Run locally

`docker run -p 8443:8443 --link ip2proxy:ip2proxy-db -t -i go-ip2proxy-api-full`

and then visit in your browser

* https://localhost:8443/

