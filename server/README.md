# Docker

## Build

```bash
docker build -t d2-roll-viewer-server .
```

## Run

```bash
docker run -d -it -p 8080:8080 --name=d2rvs d2-roll-viewer-server
```

## Attach

```bash
docker attach d2rvs
```

## Remove container 

```bash
docker container rm d2rvs
```