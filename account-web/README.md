# Account Service

This is the Account service

Generated with

```
micro new ./openBank/account/account-web --namespace=openBank --alias=account --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: openBank.web.account
- Type: web
- Alias: account

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
./account-web
```

Build a docker image
```
make docker
```