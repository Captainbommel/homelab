# Self Signed Certificates for https on the local network


## Generate a Self-Signed SSL Certificate:

```bash
sudo mkdir /etc/ssl/certs/mycert
cd /etc/ssl/certs/mycert

# Generate a private key
sudo openssl genpkey -algorithm RSA -out privkey.pem

# Generate a certificate signing request
sudo openssl req -new -key privkey.pem -out cert.csr

# Generate the self-signed certificate
sudo openssl x509 -req -days 365 -in cert.csr -signkey privkey.pem -out cert.pem
```

## Edit Nginx config to Use HTTPS

```bash
server {
    listen 443 ssl;
    server_name raspberrypi.local;

    ssl_certificate /etc/ssl/certs/mycert/cert.pem;
    ssl_certificate_key /etc/ssl/certs/mycert/privkey.pem;

    location / {
        proxy_pass http://raspberrypi.local:xxxx;  # Change port as needed
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}

server {
    listen 80;
    server_name raspberrypi.local;
    
    # Redirect HTTP to HTTPS
    return 301 https://$host$request_uri;
}

```