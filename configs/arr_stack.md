# Arr stack setup (without VPN)

Deploy ARR apps using a single command, including Jellyfin and qBittorrent.

### Useful Links
- [Servarr Wiki](https://wiki.servarr.com/)
- [Trash Guides](https://trash-guides.info/)
- [Ascii ART](https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow)

### Instructions

1. **Set Permissions**
    Go to the folder specified by `ARRPATH` and run:
    ```sh
    chown -R 1000:1000 Arr
    ```
    Now you can log in and configure all services.

2. **qBittorrent**
    - View logs:
      ```sh
      docker logs qbittorrent
      ```
    - Configure WebUI:
      - Go to Tools > Options > WebUI
      - Change user and password
      - Tick 'bypass authentication for clients on localhost'

3. **Prowlarr**
    - Access: `http://serverip:9696`
    - Settings > Download Clients > `+` > Add download client > Choose qBittorrent

4. **Sonarr**
    - Access: `http://localhost:8989`
    - Settings > Media Management > Add Root Folder > Set `/data/tvshows`
    - Settings > Download Clients > `+` > Choose qBittorrent
    - Settings > General > API key > Copy
    - Prowlarr > Settings > Apps > `+` > Sonarr
    - Settings > General > Show advanced > Set `/data/Backup` folder

5. **Radarr**
    - Access: `http://localhost:7878`
    - Settings > Media Management > Add Root Folder > Set `/data/movies`
    - Settings > Download Clients > `+` > Choose qBittorrent
    - Settings > General > API key > Copy
    - Prowlarr > Settings > Apps > `+` > Radarr
    - Settings > General > Show advanced > Backups > Set `/data/Backup` folder

6. **Lidarr**
    - Access: `http://localhost:8686`
    - Follow the same setup process as Sonarr and Radarr

7. **Readarr**
    - Access: `http://localhost:8787`
    - Follow the same setup process as Sonarr and Radarr

8. **Prowlarr Indexers**
    - Prowlarr > Indexers > `Add indexer` > Search for indexer (e.g., 'rarbg', 'yts') > Test > Save
    - Click 'Sync App Indexers' icon
    - Settings > Apps > Ensure 'Full sync' is green for each application

9. **Jellyfin**
    - Access: `http://localhost:8096`
    - Add media libraries matching folders in `docker-compose.yml`:
      - `/data/movies`
      - `/data/tvshows`
      - `/data/music`
      - `/data/books`

ARR stack setup is complete. You can now add movies in Radarr or series in Sonarr and trigger the download process.

