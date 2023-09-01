## Software development, frontend and backend
- basically software is composed by:
    - frontend: what user see
    - backend: logic and data processing
- client server architecture
    - server is remote
    - cliente is local 
    - client make requests, and server response
- backend: is focused in databases and bussiness logic
    - send, store, eliminate, modify data
    - parts:
        - server: machine or app like NGINX that recieves requests
        - application: what manages logic
        - database: used to store data
    - resposabilities of backend devs:
        - efficiente transfer of data to frontend
        - optimization and security
        - handle bug and errors
        - cloud hosting solutions and CI/CD pipelines
- APIs
    - components:
        - technical specifications: protocols and options of data exchange
        - the code: code that makes work the specification
    - Types:
        RPC (remote procedure call)
        SOAP (simple access object protocol)
        RESTFUL (request specification transfer) use http rules:
            1. The URL is a resource identifier
            2. HTTP verbs are identifiers of operations
            3. HTTP responses are representations of resources
            4. Links are relations between resources
            5. A parameter is an authentication token
- REST APIS:
    - http methods:
        - get: obtain data
        - post: send data
        - delete: eliminate data
        - put: modify data (needs all the fields)
        - head: query only metadata
        - options: debugging, used to know active web servers
        - patch: applies partially changes
        - trace: trace the path of requests
        - connect: request server as a proxy
- Django
    - Uses MVC (model view controller) architecture
        - Model: data related logic and methods of CRUD operations
        - View: handles UI logic of app
        - Controler: layer between model and view, manipulates data (i/o) and creates the render
    - In django is MVT (model view template)
        - Model = Model
        - View = Controller
        - Template = View
- DRF (django rest framework)
    - Makes django an API
    - Replace view with routes and endpoints

- Working w/ django:
    - create an enviroment w/ pip3
    - install django: echo Django==4.0 > requeriments.txt and pip3 install -r requirements.txt
    - make migrations: propagate changes made to the model in db schemes
- Elements in project   
    - manage.py utility to run commands to manage django
    - Project directory:
        - urls.py: urls to access to resources in the project
        - wsgi.py: for deployment
        - asgi.py: asynchronous support
        - settings.py: configurations of the project
- Installing postgresql
    - the site provides the necessary commands
    - install de adapter for postgres to django psycopg
- creating the database

## What is Django
- Most famous backend fw in python
