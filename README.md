# goTest
Fun with go




# Helpful docker commands
###### Build Everything
```
docker-compose up -d
```

##### Destroy Everything
```
docker-compose down
```

##### Build all docker images
```
docker-compose build
```

##### Kill docker container
```
docker kill <container_name>
```

##### Restart docker container
```
docker restart <container_name>
```

##### Get (all) docker containers
```
docker ps -a
```

##### Get all docker images
```
docker images
```

##### Remove docker container
```
docker rm <container_name>
```

##### Remove docker image
```
docker rm <image_name>
```




# Opening up the app
After running docker-compose up. You should be able to find the app at <docker-machine ip>:80

This will have the nginx reverse proxy set up to serve static assets in front of the go webserver
