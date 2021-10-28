# xVaultCTL (CLI)

The `xvaultctl` is used to control vaults in `Google Cloud`, `AWS Cloud` and `HashCorp`

## Downloading dependencies
```
$ go get golang.org/x/term
```
## Build a new app
```
$ go build -o xvaultctl main.go
```
## Basic Commands
```
$ xvaultctl help

```
## Login commands

### Login on vault provider using nominal user
```
$ xvaultctl login -u <myuser> provider <aws | gcp | hashcorp> environment <dev | qa | prod>
```
### Login on vault provider using nominal user and password
```
$ xvaultctl login -u <myuser> -p '<mypass>' provider <aws | gcp | hashcorp> environment <dev | qa | prod>
```
### Login on vault provider using service user
```
$ xvaultctl login -U <service user> vaultid <vault name id> provider <aws | gcp | hashcorp> environment <dev | qa | prod>
```
### Login on vault provider using service user and passoword
```
$ xvaultctl login -U <service user> -p '<service password>' vaultid <vault name id> provider <aws | gcp | hashcorp> environment <dev | qa | prod>
```
## Contributions
Jonas Cavalcanti | <jonascavalcantineto@gmail.com>

Robson Franklin | <robsonfrankin@gmail.com>

## Community, discussion, contribution, and support
(#casterlyrock)


