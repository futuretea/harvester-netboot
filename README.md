# harvester-netboot

Declarative definition of multi-cluster Harvester network boot server.

## Features
- [x] Declarative definition
- [x] Multi-cluster support
- [ ] Webhook support
- [ ] EFI support

## How to use

1. Prepare a clean Linux server with docker and docker-compose installed.
2. Clone this repo
3. Copy the example config file.
```bash
cp config.example.yaml config.yaml
```
4. Edit the config file for your environment.
```bash
vim config.yaml
```
5. Run the docker-compose file.
```bash
docker-compose up -d
```
