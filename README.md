# sysctl-write-docker

Write kernel parameters using sysctl with Docker.

# Usage

The input to the tool is the environment variable `SYSCTL`, which is
just a JSON mapping of kernel parameters to values.

```
docker run --privileged -e 'SYSCTL={"fs.aio-max-nr":"500000"}' quay.io/dollarshaveclub/sysctl-write-docker:latest
```

# Block

Use the `SYSCTL_BLOCK=true` environment variable to have the command
block indefinitely.
