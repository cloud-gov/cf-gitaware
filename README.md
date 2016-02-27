# Cloud Foundry Git-Aware

Plugin to make the Cloud Foundry CLI include Git metadata when deploying.

## Installation

```bash
go build
cf install-plugin cf-gitaware
```

## Usage

```bash
cf git-push <args>
```

Accepts all the normal arguments as a normal `cf push`.
