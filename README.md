# web-crawler

This is simple simulation for concurrent web crawler with dumy web pages

## Setup and running the crawler

Must have golang installed version >= 12.0.0

make file consists of 2 steps: build, run
you can run all of them 

```bash
make all
```


# Build and run docker image

```bash
docker build --tag web-crawler .
```

```bash
docker run web-crawler
```
