#!/bin/bash

docker compose --env-file .env -f compose.yaml -f compose.prod.yaml up -d --build

