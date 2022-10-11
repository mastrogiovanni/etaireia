# Development

In order to run in dev mode you need to run both services and then place an nginx in front of them.

## Backend

Open a terminal and run the backend:

```
cd digital-signature-backend
go run src/main.go
```

Backend will be listening on default port 3001 

## Frontend

Open another terminal and run the frontend

```
cd digital-signature
npm run dev
```

Frontend will be listening on default port 3000 

## Load Balancer

Run Nginx load balancer

```
docker run -p 3002:80 -v $(pwd)/local-dev/nginx.conf:/etc/nginx/nginx.conf:ro nginx
```

Application will be listening on port 3002