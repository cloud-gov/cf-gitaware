# Cloud Foundry Git-Aware

Plugin to make the Cloud Foundry CLI include Git metadata when deploying.

## Installation

Requires Cloud Foundry CLI >= v6.7.0.

```bash
go build
cf install-plugin cf-gitaware
```

## Usage

```bash
cf git-push <args>
```

Accepts all the normal arguments as a normal `cf push`.

## Development

This project uses [Godep](https://github.com/tools/godep) to manage its dependencies.

```bash
go get -u github.com/tools/godep github.com/18F/cf-gitaware
cd $GOPATH/src/github.com/18F/cf-gitaware
godep restore
```

See [Developing cf CLI Plugins](https://docs.cloudfoundry.org/cf-cli/develop-cli-plugins.html) for more information.
