# sd-exam2

##Examen 2
### Alejandro Barrera Lozano
### Curso: Sistemas distribuidos

Para la ejecucion de la aplicacion debe tener el servicio de docker corriendo en su pc y docker-compose instalado

`git clone https://github.com/NonSenseGuy/sd-exam2 `

+ Procure tener los puertos 8080, 80 y 5432
+ Genere sus certificados ssl para levantar el reverse-proxy entrando al directorio nginx y ejecutando el script generate_keys.sh y dando le permisos de lectura y escritura al directorio ssl generado
+ `cd nginx`
+ `mkdir ssl && ./generate_keys.sh`
+ `chmod +rwx ssl/*
+ ejecute `docker-compose up --build`
