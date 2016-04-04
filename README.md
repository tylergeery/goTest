# goTest
Fun with go




# Helpful docker commands
Build Everything
`docker-compose up -d`

Destroy Everything
`docker-compose down`

Build docker image
`docker build -t gofun .`

Start docker container && nginx server
`docker run -d -t -p 8000:80 --name gofunserver gofun nginx -g 'daemon off;'`

Kill docker container
`docker kill gofunserver`

Restart docker container
`docker start gofunserver`

Get (all) docker containers
`docker ps -a`

Remove docker container
`docker rm gofunserver`
