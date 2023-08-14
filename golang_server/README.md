# sandbox


A general repo for testing and depositing early ideas



## Docker Container commands
must have rancher desktop running or something equivalent. 


### build container
make sure you are in the dir with the Dockerfile or else the . wont find the dockerfile
1. docker build . -t golang_server:latest




### run container
you need to expose the port with 8080:8080 to be able to reach into the container with API commands
1. docker run -p 8080:8080 golang_server:latest 
