# web-crawler

This is simple simulation for web crawler with dumy web pages

## Setup and running the crawler

Must have golang installed version >= 12.0.0

make file consists of 2 steps: build, run
you can run all of them 

```bash
make all
```


# Build docker image

```bash
docker build --tag geolocation-service .
```

```bash
docker run -p 9090:9090 geolocation-service
```
