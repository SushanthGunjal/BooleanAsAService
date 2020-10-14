# Boolean As a Service

* A GO Api for booleans
* A service which can be used to create, delete and update boolean values.


## Tech stack used 


* Golang
* Mysql
    * Gorm as orm library
* Docker

Edit the .env file according to your requirements.

## Installation

There are 3 ways to install the service

### 1. On local machine

Run the following commands
```
git clone https://github.com/SushanthGunjal/BooleanAsAService.git
cd BooleanAsAService
```

##
### Build

```
go build .
```

### Run
```
./BooleanAsAService
```

### 2. Using Docker (Recommended)

Run the following commands

```
git clone https://github.com/SushanthGunjal/BooleanAsAService.git
cd BooleanAsAService
```

### Build
```
docker build . -t boolean_as_a_service
```

### Run
```
docker run -p 8080:8080 boolean_as_a_service 
```


### 3. Using Docker image
Pull the image from dockerhub by executing the following command
```
docker image pull sushanthgunjal/boolean_as_a_service 
```

### Run
```
docker run -p 8080:8080 boolean_as_a_service  
```



## API

### BaseUrl
```
http://localhost:8080
```

### Creating a new Boolean Service:



```
request:

{
  "value":true,
   "key": "name" 
}

response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}

```

Usage 
```
curl -i -X POST http://localhost:8080 -d '{"value": true, "key": "name"}'

```

### Getting a Boolean Service:

```
GET /:id
response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}

```
Usage 
```
curl -i -X GET http://localhost:8080/[id]

```

### Updating a Boolean Service:



```
PATCH /:id
request:

{
  "value":false,
  "key": "new name" 
}
response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": false,
  "key": "new name"
}

```

Usage 
```
curl -i -X PATCH http://localhost:8080/[id] -d '{"value": false, "key": "new name"}'

```

### Deleting a Boolean Service:
```
request:
DELETE /:id
response:
HTTP 204 No Content

```

Usage 
```
curl -i -X DELETE http://localhost:8080/[id]

```
