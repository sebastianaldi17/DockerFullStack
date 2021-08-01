# DockerFullStack - Server
This folder serves as the backend for the service. The project structure is broken down into:

- Data: all data fetching or CRUD operation goes here.
- Business: all logic and checking goes here.
- Handlers: all handler function goes here.
- Entity: all structs and constants goes here.

The contents of each folder then contains multiple folders for easier separation if needed.

How it works is that all routes and cron jobs will be initialized on `main.go`.
All handlers will read the body of the request, then send the information to the business layer.
The business layer would then do the checking or logic (it may call to multiple data functions if needed), and return the fetched data back to the handler.

TLDR: handler -> business -> data. Handlers must not jump to data layer. Entity can be used anywhere.
## TODO
- Write unit tests and mocks
- Separate handlers to cron - http - ?? if needed
- Find a way for better error logging