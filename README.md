# Sistema de Gestión de Comercio Electrónico en Go

Este proyecto es un sistema básico de gestión de comercio electrónico desarrollado en el lenguaje Go. 

Incluye funcionalidades clave como:

- Registro de usuarios
- Inicio y cierre de sesión
- Catálogo de productos
- Cifrado de contraseñas con bcrypt
- Manejo de sesiones usando cookies
- Plantillas HTML renderizadas con html/template

Estructura del proyecto

```
ecommerce/
├── main.go
├── go.mod
├── models/
│   ├── user.go
│   ├── product.go
│   └── cart.go
├── utils/
│   ├── user_store.go
│   ├── session.go
│   └── product_loader.go
├── handlers/
│   ├── auth_handler.go
│   └── product_handler.go
├── templates/
│   ├── login.html
│   ├── register.html
│   └── products.html
├── data/
│   ├── users.json
│   └── products.json
```

---

Instalación de dependencias

Desde la carpeta raíz del proyecto, ejecuta:

```bash
go mod init ecommerce
go get github.com/google/uuid
go get golang.org/x/crypto/bcrypt
```

---

Cómo ejecutar el proyecto

1. Abre una terminal en la carpeta `ecommerce/`.
2. Ejecuta:

```bash
go run main.go
```

3. Abre tu navegador y visita:

- `http://localhost:8080/register` → Registro de usuario
- `http://localhost:8080/login` → Inicio de sesión
- `http://localhost:8080/catalogo` → Catálogo de productos (requiere sesión)
- `http://localhost:8080/logout` → Cierre de sesión

---

Seguridad implementada

- Las contraseñas se almacenan de forma cifrada usando `bcrypt`.
- El acceso al catálogo está protegido por sesión mediante cookies HTTP.

---

Estado actual del sistema

- [x] Registro y login funcional
- [x] Sesiones con cookies
- [x] Carga de productos desde JSON
- [x] Plantillas HTML
- [ ] Carrito de compras (próxima fase)
- [ ] Proceso de pago simulado
- [ ] Panel administrativo con métricas

---

Autor

**Santiago Ríos**  
Estudiante de Ingeniería en Sistemas de Información  
Proyecto académico - 2025
