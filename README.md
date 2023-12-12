# ProyectoFinal_Back_Go
README

Este repositorio contiene un servidor en Go (Golang) que implementa una API básica para la gestión de usuarios, incluyendo funciones para crear usuarios, iniciar sesión, consultar usuarios y eliminar usuarios. La aplicación utiliza el framework Gin para facilitar el enrutamiento y la gestión de solicitudes HTTP, y se integra con una base de datos PostgreSQL a través de GORM para el modelado y la manipulación de datos.
Requisitos previos

Asegúrate de tener instalado Go en tu sistema. Puedes descargar Go desde https://golang.org/dl/.

Además, necesitarás una base de datos PostgreSQL. Puedes instalar PostgreSQL desde https://www.postgresql.org/download/.
Configuración de la base de datos

Antes de ejecutar la aplicación, asegúrate de crear una base de datos en PostgreSQL y proporcionar las credenciales correctas en la función init del archivo main.go. Puedes ajustar la cadena de conexión en la siguiente línea:


    db, err = gorm.Open("postgres", "user=postgres password=Tu_Password_de_la_base_de_datos dbname=el_nombre_dela_basededatos port=5432 sslmode=disable")

Asegúrate de reemplazar user, password, dbname y otros parámetros según tu configuración de PostgreSQL.
Instalación de dependencias

Este proyecto utiliza varias bibliotecas externas. Puedes instalar estas dependencias utilizando el siguiente comando:

bash

    go get -u github.com/gin-gonic/gin <!-- Gin framework -->
    go get -u github.com/gin-contrib/cors <!-- Gin middleware for CORS -->
    go get -u github.com/jinzhu/gorm <!-- ORM -->
    go get -u github.com/lib/pq <!--PostgreSQL driver -->
    go get -u golang.org/x/crypto/bcrypt <!-- Password hashing -->


Ejecución de la aplicación

Una vez que hayas configurado la base de datos y hayas instalado las dependencias, puedes ejecutar la aplicación con el siguiente comando:

bash

    go run main.go
    La aplicación se ejecutará en http://127.0.0.1:8080 y aceptará solicitudes desde http://127.0.0.1:5173 gracias a la configuración CORS.

Uso de la API

La API proporciona las siguientes rutas:

    POST /crear-usuario: Crea un nuevo usuario.
    POST /iniciar-sesion: Inicia sesión con un usuario existente.
    GET /consultar-usuario/:email: Consulta la información de un usuario por correo electrónico.
    DELETE /eliminar-usuario/:id: Elimina un usuario por ID.

Consulta el código fuente para obtener detalles sobre la estructura de los datos y los endpoints.

Contribuciones

Si encuentras problemas o mejoras posibles, siéntete libre de abrir un problema o enviar una solicitud de extracción. ¡Las contribuciones son bienvenidas!

¡Disfruta usando la aplicación!
