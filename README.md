# README

Create a `.env` file with the following:

```
DATABASE_NAME=name
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=admin
```

# Run in Dev mode

Run the application (dev mode) with the following:

```
docker-compose --env-file .env --profile dev up --build
```

Access the app at `http://localhost`

# Run in Prod mode

Run the application (prod mode) with the following:

```
docker-compose --env-file .env --profile prod up --build
```

Access the app at `https://${NODE_HOST}`

where `NODE_HOST` is the domain specified in the `.env` file

# Code for mobile phone

If your website is exposed to internet you can create a QRcode for it:

```
npm install -g qrcode
```

```
qrcode https://$(curl 'https://api.ipify.org')
```
