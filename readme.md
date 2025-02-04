# Homelab Setup

<br>

## Base Stack 

Services will connect Nginx Proxy Manager over the internal network, without opening any ports.

- Nginx Proxy Manager - encrypted outward traffic
- Uptime Kuma - monitoring
- Portainer - container management
- Gotify - notifications 

## Connected Services
- Watchtower - container updates
- Wallos - subscription tracking
- Vaultwarden - password manager


```yml
# the services will be connected to the base stack network
# in the docker compose file like this:

networks:
      - network-example

networks:
  network-example:
    external: true
```

<br>

## Local Services
- SearXNG
- LibreSpeed (to test wifi speeds)

<br>

## DynDNS

[Cloudflare Documentation](https://developers.cloudflare.com/api/node/resources/dns/subresources/records/methods/export/)

[Fritz!Box Documentation](https://avm.de/service/wissensdatenbank/dok/FRITZ-Box-7590/30_Dynamic-DNS-in-FRITZ-Box-einrichten/)

[Github: favonia/cloudflare-ddns](https://github.com/favonia/cloudflare-ddns)
