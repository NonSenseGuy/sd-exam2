# Examen 2
### Alejandro Barrera Lozano
### Curso: Sistemas distribuidos

Golang API que usa Postgres como base de datos, la api consiste en un inventario de personas reportadas por fraude.

El modelo consisten en 

FraudataItem
+ name	string
+ is_reported	bool
+ report_reasons   string
+ created_on	time
+ updated_on	time

Para la ejecucion de la aplicacion debe tener el servicio de docker corriendo en su pc y docker-compose instalado

`git clone https://github.com/NonSenseGuy/sd-exam2 `

+ Procure tener los puertos 8080, 80 y 5432 abiertos
+ Genere sus certificados ssl para levantar el reverse-proxy entrando al directorio nginx y ejecutando el script generate_keys.sh y dandole permisos de lectura y escritura al directorio ssl generado
`cd nginx`
`mkdir ssl && ./generate_keys.sh`
`chmod +rwx ssl/*`
+ ejecute `docker-compose up --build`

Para el consumo de la api haga las peticiones al localhost/api/v1 a continuacion unos ejemplos,

### Conocer la salud del servicio 

GET /api/v1/health

### Listar todos los casos reportados
Maximo se pueden listar 200 reportes
GET /api/v1/fraudata

### Obtener un reporte por su id
GET /api/v1/fraudata/item?id=x

### Eliminar un reporte por su id
DELETE /api/v1/fraudata/item?id=x


### Agregar un nuevo reporte
POST /api/v1/fraudata 

Request Body:
`{
	"name": "Alejandro Barrera",
	"is_reported": true,
	"report_reasons": "identity thief"
}`
