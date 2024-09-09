## Configuración del Proyecto

1. Clonar el repositorio:
   git clone https://github.com/usuario/talentpitch_api.git
   cd talentpitch_api  `

2. Inicializar el módulo de Go:
   go mod tidy 

3. Instalar MySQL en tu sistema operativo si no está instalado, y configurar las credenciales. Puedes hacerlo con los siguientes comandos:

   sudo apt-get install mysql-server
   sudo mysql_secure_installation
   
4. Iniciar MySQL y crear la base de datos:

   mysql -u root -p
   CREATE DATABASE talentpitch_db;
   CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
   GRANT ALL PRIVILEGES ON talentpitch_db.* TO 'user'@'localhost';
   FLUSH PRIVILEGES;

5. Actualizar el archivo .env

Asegúrate de que el archivo .env esté configurado para usar MySQL:

   DB_HOST=localhost
   DB_USER=user
   DB_PASSWORD=password
   DB_NAME=talentpitch_db
   DB_PORT=3306

5. Ejecutar la aplicación:
   go run cmd/main.go

## Correr el Proyecto con Docker-Compose

1. Asegurarse de tener Docker y Docker-Compose instalados.

2. Clonar el repositorio:
   git clone https://github.com/usuario/talentpitch_api.git
   cd talentpitch_api

3. Construir y ejecutar los servicios:
   
   docker-compose up --build

4. Acceder a la API:
   La aplicación estará disponible en http://localhost:8080.
