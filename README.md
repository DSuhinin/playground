# Test Task Service.

# Topics
* [Common Info](#common-info)
* [Prerequisites](#Prerequisites)
* [Set up and run service](#set-up-and-run-service)
* [Environment Variables](#environment-variables)
* [Endpoints](#endpoints) 
  * [Sequence Endpoints](#sequence-endpoints)
    * [POST /sequences](#post-sequences)
  * [Step Endpoints](#step-endpoints)
    * [PUT /steps/{step_id}](#put-stepsstep_id)
    * [DELETE /steps/{step_id}](#delete-stepsstep_id)
    * [PUT /steps/{step_id}/properties/open-tracking-enabled](#put-stepsstep_idpropertiesopen-tracking-enabled)
    * [PUT /steps/{step_id}/properties/click-tracking-enabled](#put-stepsstep_idpropertiesclick-tracking-enabled)
  * [Service Endpoints](#service-endpoints)
    * [POST /service/status](#get-serviceinfo)
    * [GET /service/info](#get-servicestatus)
* [Appendix A. Response codes](#appendix-a-response-codes)
* [Appendix B. Request headers](#appendix-b-request-headers)

# Prerequisites
- Golang
- Postgres
- Docker
- Make

# Set up and run service

- go to the project folder root.
- run `make docker_build_image service_start` to build and run everything in Docker.
- run `make lint` to run linter over the code.
- run `make go_test_unit` to run `unit` tests.
- run `make docker_build_image service_test` to run `integration` tests.

# Environment Variables

| Name          		 | Required | Example		| Description	| 
|-----------| --------- | ----------------- | ------------- |		 
| POSTGRES_USER     		 | Y         | root 		| Database user.|
| POSTGRES_PASS 			 | Y         | password    	| Database password.|
| POSTGRES_HOST 			 | Y         | localhost 	| Database host.|
| POSTGRES_PORT 			 | Y         | 3306 		| Database port.|
| POSTGRES_DB_NAME 		 | Y         | database_name 	| Database name.|
| SERVER_ADDRESS 		 | N         | 127.0.0.1:8080 	| Server address and port where service is listening for incoming requests. If not provided then `127.0.0.1:8080` will be used.|
| LOG_LEVEL 			 | Y         | debug 		| Logger level. Possible values are: debug, info, warn, error.|

# Endpoints
# Sequence Endpoints
## POST /sequences
Endpoint create new `sequence` object.

**Request info**

```
HTTP Request method    POST
Request URL            https://localhost:8080/sequences
```

**Request body**

```
{
    "name": "Sequence Name",
	"steps": [{
	  "subject": "Step Subject",
	  "content": "Step Content"
	}]
}
```

**Response body**

```
{
    "id": 1,
	"name": "Sequence Name",
	"steps": [{
	  "id": 1,
	  "subject": "Step Subject",
	  "content": "Step Content
	  "open_tracing_enabled: false,
	  "click_tracing_enabled: false,
	  "created_at": "2024-04-21T14:09:45.955366+02:00",
	  "updated_at": "2024-04-21T14:09:45.955366+02:00",
	}],
	"created_at": "2024-04-21T14:09:45.955366+02:00",
	"updated_at": "2024-04-21T14:09:45.955366+02:00",
}
```


# Step Endpoints
## PUT /steps/{step_id}
Endpoint to update existing `step` data.

**Request info**

```
HTTP Request method    PUT
Request URL            https://localhost:8080/steps/{step_id}
```
**Request body**

```
{
  "subject": "Step Subject",
  "content": "Step Content"
}
```

**Response body**

```
{
  "id": 1,
  "subject": "Step Subject",
  "content": "Step Content",
  "open_tracing_enabled": false,
  "click_tracing_enabled": false,
  "created_at": "2024-04-21T14:09:45.955366+02:00",
  "updated_at": "2024-04-21T14:09:45.955366+02:00"
}
```

## PUT /steps/{step_id}/properties/open-tracking-enabled
Endpoint to update `open_tracking_enabled` property of  existing `step`.

**Request info**

```
HTTP Request method    PUT
Request URL            https://localhost:8080/steps/{step_id}/properties/open-tracking-enabled
```
**Request body**

```
{
  "state": true
}
```

**Response body**

```
{
  "id": 1,
  "subject": "Step Subject",
  "content": "Step Content",
  "open_tracing_enabled": true,
  "click_tracing_enabled": false,
  "created_at": "2024-04-21T14:09:45.955366+02:00",
  "updated_at": "2024-04-21T14:09:45.955366+02:00"
}
```

## PUT /steps/{step_id}/properties/click-tracking-enabled
Endpoint to update `click_tracking_enabled` property of  existing `step`.

**Request info**

```
HTTP Request method    PUT
Request URL            https://localhost:8080/steps/{step_id}/properties/click-tracking-enabled
```
**Request body**

```
{
  "state": true
}
```

**Response body**

```
{
  "id": 1,
  "subject": "Step Subject",
  "content": "Step Content",
  "open_tracing_enabled": false,
  "click_tracing_enabled": true,
  "created_at": "2024-04-21T14:09:45.955366+02:00",
  "updated_at": "2024-04-21T14:09:45.955366+02:00"
}
```

## DELETE /steps/{step_id}
Endpoint to delete existing `step` data.

**Request info**

```
HTTP Request method    DELETE
Request URL            https://localhost:8080/steps/{step_id}
```
**Request body**

```
{}
```

**Response body**

```
{}
```

# Service Endpoints
## GET /service/info
The endpoint to get service info.

**Request info**

```
HTTP Request method    GET
Request URL            https://localhost:8080/service/info
```
**Request body**

```
{}
```

**Response body**

```
{
  "dependencies": {
    "postgres": {
      "status": 200,
      "latency": 0.002757021
    }
  }
}
```
**Note**
Endpoint give the full picture about `service` current state. Endpoint mostly needs for Kubernetes readiness probe.

## GET /service/status
The endpoint to get service status.

**Request info**

```
HTTP Request method    GET
Request URL            https://localhost:8080/service/status
```

**Request body**

```
{}
```

**Response body**
```
{}
```
**Note**
Endpoint give the full picture about `service` current state. Endpoint mostly needs for Kubernetes liveness probe.

# Appendix A. Response codes

### HTTP error codes

Application uses standard HTTP response codes:
```
200 - Success
201 - Created
400 - Request error
401 - Authentication error
404 - Entity not found
500 - Internal Server error
```

Additional information about the error is returned as JSON-object like:
```
{
    "code": "{error-code}",
    "message": "{error-message}"
}
```
# Appendix B. Request headers
HTTP header required for API calls:
- `Authorization` : `Bearer token` - authorization based on `JWT` token.
