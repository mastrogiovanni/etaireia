# Etaireia - a simple solution for Digital Signature for NPO, caritative organization and small companies. 

Etaireia is a software to allow to sign documents via mobile app and desktop leveraging 
a central point where documents are stored, and keps secure with minimal maintenance effort.

Everything is work in progress. 

We will update this document a lot of times in this stage.

## Development Guidelines

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

Deploy the server in a public place and add DNS information to `.env`:

```
# For example...
NODE_HOST=etaireia.xxx.yyy
```

Run the application (prod mode) with the following:

```
docker-compose --env-file .env --profile prod up --build
```

Access the app at `https://${NODE_HOST}`

where `NODE_HOST` is the domain specified in the `.env` file

