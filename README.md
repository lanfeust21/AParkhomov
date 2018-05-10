
# Alexander Parkhomov Tables

The original file given by **Alexander Parkhomov** is located in **db/FusFis.zip**

The content has been converted in csv files located in db/csv.
The csv files has been imported in a _sqlite3 database_ located in db/sqlite/aparkhomov.db3

You can analyse the data with sql and the interface directly with the application found on http://sqlitebrowser.org/ for each system.

#1 to run the application with go (for developper)

The application will run a webserver locally on the port 8080 by default. It is developped in Golang and use the beego framework with a sqlite3 driver. 

## setup go
go to https://golang.org/dl/ and download the installer 

## setup beego and bee (at the console)
>go get github.com/beego/bee

## start local webserver (at the console)
> cd webapp
> bee run 

open the browser and enter http://localhost:8080/


# Docker

The DockerFile allow to run the application in a docker container. it allows you to run it locally with docker installed or on any server with docker installed.




