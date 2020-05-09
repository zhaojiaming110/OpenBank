# Timing Service

This is the Timing service

Generated with

```
micro new timing-server --namespace=micro.open.bank --alias=timing --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: micro.open.bank.service.timing
- Type: service
- Alias: timing

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./timing-service
```

Build a docker image
```
make docker
```