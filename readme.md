# Homelab Setup

A collection of Docker Compose files for self-hosted services, organized as modular stacks that can be deployed independently or together.

## Base Stack

Core infrastructure services defined in `compose-files/base-stack.yaml`. Other services connect over the internal Docker network without exposing ports.

| Service | Description |
|---------|-------------|
| [Cloudflared](https://github.com/cloudflare/cloudflared) | Cloudflare tunnel for secure ingress |
| [Portainer](https://www.portainer.io/) | Docker container management UI |
| [Uptime Kuma](https://github.com/louislam/uptime-kuma) | Service monitoring and status pages |
| [Gotify](https://gotify.net/) | Push notification server |
| [Mediatracker](https://github.com/bonukai/MediaTracker) | Track movies, TV shows, books, and games |
| [Wallos](https://github.com/ellite/Wallos) | Subscription and expense tracker |

## Networking & DNS

| Service | File | Description |
|---------|------|-------------|
| [Nginx Proxy Manager](https://nginxproxymanager.com/) | `npm.yaml` | Reverse proxy with SSL management |
| [Cloudflare DDNS](https://github.com/favonia/cloudflare-ddns) | `cloudflare-ddns.yaml` | Automatic DNS record updates with current IP |
| [Cloudflared](https://github.com/cloudflare/cloudflared) | `cloudflared.yaml` | Standalone Cloudflare tunnel |
| [OpenVPN](https://openvpn.net/) | `openvpn.yaml` | VPN Access Server |

## Media & Downloads

| Service | File | Description |
|---------|------|-------------|
| [Navidrome](https://www.navidrome.org/) | `navidrome.yaml` | Music streaming server (Subsonic-compatible) |
| [Pinchflat](https://github.com/kieraneglin/pinchflat) | `pinchflat.yaml` | YouTube media downloader and organizer |

## Productivity & Files

| Service | File | Description |
|---------|------|-------------|
| [Mealie](https://mealie.io/) | `mealy.yaml` | Recipe management and meal planning |
| [Wiki.js](https://js.wiki/) | `wiki.yaml` | Wiki engine with PostgreSQL backend |
| [Seafile](https://www.seafile.com/) | `seafile.yaml` | File sync (MariaDB + Memcached backend) |
| [Vaultwarden](https://github.com/dani-garcia/vaultwarden) | `vaultwarden.yaml` | Bitwarden-compatible password manager |

## Monitoring & Automation

| Service | File | Description |
|---------|------|-------------|
| [Watchtower](https://containrrr.dev/watchtower/) | `watchtower.yaml` | Automated container updates with Gotify notifications |
| [Changedetection.io](https://github.com/dgtlmoon/changedetection.io) | `changedetection.yaml` | Website change monitoring |
| [OliveTin](https://www.olivetin.app/) | `olivetin.yaml` | Web UI for running predefined shell commands |

## Dashboards

| Service | File | Description |
|---------|------|-------------|
| [Homepage](https://gethomepage.dev/) | `homepage.yaml` | Dashboard with Docker integration |
| [Dashy](https://dashy.to/) | `dashy.yaml` | Customizable dashboard |

## Privacy & Tools

| Service | File | Description |
|---------|------|-------------|
| [SearXNG](https://docs.searxng.org/) | `searxng.yaml` | Privacy-respecting metasearch engine (with Redis) |
| [LibreSpeed](https://github.com/librespeed/speedtest) | `librespeed.yaml` | Self-hosted network speed test |

## Other

| Service | File | Description |
|---------|------|-------------|
| [Dumbware](https://dumbware.io/) | `dumb_stack.yaml` | Dumbdo (tasks), Dumbkan (kanban), Dumbdrop (file sharing) |
| [Minecraft](https://docker-minecraft-server.readthedocs.io/) | `minecraft.yaml` | Minecraft game server |

## Connecting Services to the Base Stack

Services that need to reach the base stack network can join it as an external network:

```yml
networks:
  homelab_default:
    external: true
```

## Environment Variables

Most compose files use environment variables for sensitive or host-specific values. Create a `.env` file in the compose directory with the required variables:

| Variable | Used By |
|----------|---------|
| `CF_API_TOKEN` | base-stack, cloudflared |
| `GOTIFY_URL`, `GOTIFY_TOKEN`, `GOTIFY_HOSTNAME` | watchtower |
| `MEDIATRACKER_DATA_PATH` | base-stack |
| `DB_USER`, `DB_PASS` | wiki |
| `DB_ROOT_PASSWORD`, `SEAFILE_ADMIN_EMAIL`, `SEAFILE_ADMIN_PASSWORD` | seafile |
| `DUMB_PIN`, `DUMBDROP_UPLOAD_PATH` | dumb_stack |
| `HOST` | vaultwarden |
| `CLOUDFLARE_API_TOKEN`, `DOMAINS` | cloudflare-ddns |

## Configs

The `configs/` directory contains example configuration files for:

- **Homepage** — `homepage_services.yaml`, `homepage_settings.yaml`, `homepage_widgets.yaml`
- **OliveTin** — `olivetin_config.yaml`
