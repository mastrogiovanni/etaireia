# README

Create a `.env` file with the following:

```
DATABASE_NAME=name
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=admin
```

Run the application with:

```
docker-compose --env-file .env up --build
```

# Code for mobile phone

If your website is exposed to internet you can create a QRcode for it:

```
npm install -g qrcode
```

```
qrcode https://$(curl 'https://api.ipify.org')
```
