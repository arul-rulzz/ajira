# AjiraNet Service

[![N|Solid](https://media.glassdoor.com/sql/1168244/ajira-squarelogo-1511324401993.png)](https://nodesource.com/products/nsolid)



A simple service to achieve the followings.
  - Add Devices
  - Create Connection Between Devices
  - Fetch the route between two devices
  - Fetch available devices


# Installation 
## Docker
Create the following network layers in docker. Ignore if you already created.
  - ajira_internal
  - ajira_default

Network creation command sample.
```sh
$ sudo docker network create ajira_internal
```

If you are going to run directly means, execute the 
```sh
$ sh docker_build_and_run.sh
```
shell script.

Feel free to change the envs and network layer. Need to modify the docker files, for the respective changes.

## CommandLine

To download the dependencies, execute the follwoing
```sh
$ go mod download
```

To start the service
```sh
$ sh run.sh
```

Change the env to modify the port no. 