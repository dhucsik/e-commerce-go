# E-commerce backend

So far, 2 services. Main app and email service. 

They communicate using grpc. I plan to increase the number of services in the future and also use message brokers.

## How to start

Create your app.env file in both backend and email-send. Set variables using .env.example files as example.
```
docker-compose up --build
```