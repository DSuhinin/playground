branch | build 
-------|-------
master |[![Build Status]() |


# Topics
* [Overview](#overview)
* [Dependencies](#dependencies)
* [Settings](#settings)
* [Endpoints](#endpoints)
    * [Deal Endpoints](#deal-endpoints)
        * [GET /deals](#get-deals)
        * [GET /deals/{deal_id}](#get-applicationapplication_idsummary)
* [Service Endpoints](#service-endpoints)
    * [GET /_service/status](#get-_servicestatus)
    * [GET /_service/info](#get-_serviceinfo)
* [Appendix A. Response codes](#appendix-a-response-codes)
* [Appendix B. Error codes](#appendix-b-error-codes)
    * [HTTP Status code = 400](#error-codes-with-http-status-code--400)
    * [HTTP Status code = 500](#error-codes-with-http-status-code--500)

# Overview
Service provide an API to work with Deal entity.

## Dependencies
* MySQL
* GO 1.12.5

## Settings
 Env                                             | Description
-------------------------------------------------|-------------------------------------------------------------------------------------------
 DEMO_SERVICE_SERVER_ADDRESS                     | Listening address (by default: `127.0.0.1:8080`)
 DEMO_SERVICE_SERVER_READ_TIMEOUT                | ReadTimeout is the maximum duration for reading the entire request, including the body. (by default: `5s`)
 DEMO_SERVICE_SERVER_WRITE_TIMEOUT               | WriteTimeout is the maximum duration before timing out writes of the response.
 DEMO_SERVICE_LOG_LEVEL                          | Log level (by default: `error`)
 DEMO_SERVICE_MYSQL_DSN                          | MySQL DSN, e.g: root:root@tcp(localhost:33306)/deals_manager

## Endpoints

### Deal Endpoints

#### GET /deals
Retrieves the list of Deal entities.

**Request info**

```
HTTP Request method    GET
Request URL            /deals
```

**Request**
```json
```

**Response info**
```
HTTP/1.1 200 OK
Content-Type: application/json
```

**Response**
```json
{
  "data": [
    {
      "deal_id": 761,
      "opportunity_id": "O0-761",
      "contract_id": 0,
      "kw_uid": "556396",
      "kw_uid_name": "Nicole Burton",
      "mc_id": 2,
      "mc_key": 0,
      "checklist_ids": "",
      "deal_owner": "615826",
      "deal_owner_name": "",
      "deal_name": "Buyer Deal"
    },
    {
      "deal_id": 762,
      "opportunity_id": "O0-762",
      "contract_id": 0,
      "kw_uid": "556396",
      "kw_uid_name": "Nicole Burton",
      "mc_id": 2,
      "mc_key": 0,
      "checklist_ids": "",
      "deal_owner": "556396",
      "deal_owner_name": "Nicole Burton",
      "deal_name": "Buyer Deal"
    },
    {
      "deal_id": 763,
      "opportunity_id": "O0-763",
      "contract_id": 0,
      "kw_uid": "556396",
      "kw_uid_name": "Nicole Burton",
      "mc_id": 2,
      "mc_key": 0,
      "checklist_ids": "",
      "deal_owner": "556396",
      "deal_owner_name": "Nicole Burton",
      "deal_name": "Buyer Deal"
    },
    {
      "deal_id": 764,
      "opportunity_id": "O0-764",
      "contract_id": 0,
      "kw_uid": "556396",
      "kw_uid_name": "Nicole Burton",
      "mc_id": 101,
      "mc_key": 0,
      "checklist_ids": "",
      "deal_owner": "556396",
      "deal_owner_name": "Nicole Burton",
      "deal_name": "Buyer Deal"
    }
  ]
}
```


#### GET /deals/{deal_id}
Retrieves Deal entity by Deal ID.

**Request info**

```
HTTP Request method    GET
Request URL            /deals/{deal_id}
```

**Request**
```json
```

**Response info**
```
HTTP/1.1 200 OK
Content-Type: application/json
```

**Response**
```json
{
  "deal_id": 888,
  "opportunity_id": "",
  "contract_id": 0,
  "kw_uid": "556396",
  "kw_uid_name": "Nicole Burton",
  "mc_id": 2,
  "mc_key": 0,
  "checklist_ids": "",
  "deal_owner": "556396",
  "deal_owner_name": "Nicole Burton",
  "deal_name": "Buyer Deal"
}
```

# Service Endpoints
## GET /_service/status
Internal endpoint to get service health status information. In case of good service state endpoint returns HTTP Status Code 200. In any other case endpoint returns HTTP Status Code = 400.

**Request info**

```
HTTP Request method    GET
Request URL            /_service/status
```

**Request body**

```json
{}

```

**Response info**
```
HTTP/1.1 200 OK
Content-Type: application/json
```

**Response**
```json
{}
```

## GET /_service/info
Internal endpoint to get extended service health status information. In case of good service state endpoint returns HTTP Status Code 200. In any other case endpoint returns HTTP Status Code = 400. Also endpoint returns some additional data about each dependency used in this service.

**Request info**

```
HTTP Request method    GET
Request URL            /_service/info
```

**Request body**

```json
{}

```

**Response info**
```
HTTP/1.1 200 OK
Content-Type: application/json
```

**Response**
```json
{
	"build": {
		"date": "201708111723",
		"branch": "v1",
		"commit": "980c192d5711248670835545b96ee54a2596bb77"
	},
	"dependency_list": {
		"mysql_db": {
			"status": 200,
			"latency": 0.000209081
		}
	}
}
```


## Appendix A. Response codes

### HTTP error codes

Application uses standard HTTP response codes:
```
200 - Success
400 - Request error
404 - Entity not found
405 - Method not allowed
500 - Server error
```

## Appendix B. Error Codes

#### Error codes with HTTP Status code = 400
 
 Code    |  Message
---------|-------------------------------------------------------------------------------------------
 40000   | json parsing error.
 40001   | application not found.

#### Error codes with HTTP Status code = 500

 Code    |  Message
---------|-------------------------------------------------------------------------------------------
 10000   | internal server error.
