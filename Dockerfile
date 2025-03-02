# Usar una imagen base de Go
FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /ecommerce_web

# Copiar el código fuente de tu proyecto al contenedor
COPY . .

# Compilar la aplicación Go
RUN go build -o ecommerce

# Exponer el puerto en el que se ejecuta la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]