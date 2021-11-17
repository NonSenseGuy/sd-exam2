### Alejandro Barrera Lozano
### Distributed Systems Course

## Running the app

Beforehand you should have this dependencies installed

+ docker (also be sure to be running it as a service already)
+ docker-compose 

Now you are ready to clone the repository
	`git clone https://github.com/NonSenseGuy/sd-exam2 `

Generate your ssl certs, required by the reverse proxy
	`cd nginx && ./generate_keys.sh`

Build and deploy the API
	`docker-compose up --build`

## API 

This api servers to make a database of persons reported doing fraud

The model you are going to get and post to the databse is this one

FraudataItem
+ name	string
+ is_reported	bool
+ report_reasons   string
+ created_on	time
+ updated_on	time


To consume the API you are going to do request to this address ['localhost/api/v1']('localhost/api/v1')

## API Requests

### Service Health 

GET /health

```json
{
  "health": "Running"
}
```

### List all reported cases

GET /api/v1/fraudata

```json
{
  "items": [
    {
      "id": "1637164514-0595213501-8190657423",
      "name": "Benito Martinez",
      "is_reported": true,
      "report_reasons": "Exceso de facha",
      "created_on": "2021-11-17T15:55:14.595218Z",
      "update_on": "0001-01-01T00:00:00Z"
    },
    {
      "id": "1637179892-0682314649-0621784935",
      "name": "Alejandro Barrera",
      "is_reported": false,
      "created_on": "2021-11-17T20:11:32.682329Z",
      "update_on": "0001-01-01T00:00:00Z"
    },
  ]
}
```

### Get a report by id
GET /api/v1/fraudata/item?id=x

```json
{
  "item": {
    "id": "1637164514-0595213501-8190657423",
    "name": "Benito Martinez",
    "is_reported": true,
    "report_reasons": "Exceso de facha",
    "created_on": "2021-11-17T15:55:14.595218Z",
    "update_on": "0001-01-01T00:00:00Z"
  }
}
```

### Delete a report by id
DELETE /api/v1/fraudata/item?id=x


### Create a report
POST /api/v1/fraudata 

Request Body:
`{
	"name": "Alejandro Barrera",
	"is_reported": true,
	"report_reasons": "Idk u tell me"
}`


## What's missing to deploy to a production environment?

+ Using a domain to host it
+ Use and ec2 as a server
+ Improve reliability by having replicas with kubernetes or docker swarm
+ Make script to build the app, run functional tests and deploy if the tests passed
