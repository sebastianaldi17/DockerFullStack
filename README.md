# DockerFullStack
DockerFullStack is a mini project that I am working on.
The frontend is made using react.js, while the backend is made using Go and PostgreSQL.

This repository is separated into (currently) three folders:
- dbscripts: .sql files that will be initialized on startup.
- server: Go code containing the backend logic.
- web: react.js code containing the frontend.

To run the backend only, you can do `docker compose up server`. If you want to run all services, do `docker compose up` instead.
Frontend side is currently a fresh `create-react-app` project, without any changes yet.