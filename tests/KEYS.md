openssl genrsa -out key.pem 4096

openssl rsa -in key.pem -out pubkey.pem -pubout -outform pem

openssl req -new -x509 -key key.pem -out cert.pem -days 1095
