# Chat local en terminal

Chat local construido en Go: Cliente / Sevidor

# Como usar

- Instala Go
- Descarga los archivos server.go y cliente.go
- Cambia la parte comentada en el código ip:port con el ip de tu computadora el puerto puedes dejar el mismo
- Compila los archivos en la terminal
```
  go build server.go
  go build cliente.go
```

Inicia primero el server. Te pedira permisos para abrir y utilizar los puertos. Recuerda que tienen que estar en la misma red.
Ahora puedes iniciar el cliente. Puedes iniciar multiples clientes en el mismo ordenador o en otros equipos recuerda que para esto debes poner la ip del ordenador al que te quieres conectar tanto en server.go (linea 73) como en cliente.go (linea 15)
antes de compilarlos.

Es posible conectar equipos Windows, Mac y Linux al programa al mismo tiempo. En Mac y algunas distribuciones linux es posible el uso de emojis.
