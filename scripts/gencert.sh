openssl genrsa -out private_key 2048
openssl req -new -x509 -key private_key -out cert.pem -days 365 -subj "/C=US/ST=Washington/L=Seattle/O=JJ/CN=localhost"
cp cert.pem ../cert.pem
cp private_key ../private_key