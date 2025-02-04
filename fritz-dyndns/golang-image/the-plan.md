# Updating Cloudflare DNS with Fritz!Box DynDNS

<br>

## Idea:

1. The web server listens on port 8080 and exposes an /update endpoint.
2. Fritz!Box sends DynDNS updates by making a request to http://myserver:8080/update?myip=<new_ip>.
3. The server extracts the myip parameter and forwards it to Cloudflare using the Cloudflare API.
4. Cloudflare updates the DNS record with the new IP.

## Environment Variables:

- CLOUDFLARE_ZONE_ID: Cloudflare zone ID
- CLOUDFLARE_DNS_RECORD_ID: ID of the DNS record to update
- CLOUDFLARE_API_TOKEN: Cloudflare API token
- CLOUDFLARE_HOSTNAME: The hostname