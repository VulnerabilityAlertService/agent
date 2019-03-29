# [vas] - Vulnerability diagnostic tool.

[![Build Status](https://travis-ci.org/VulnerabilityAlertService/vas.svg?branch=master)](https://travis-ci.org/VulnerabilityAlertService/vas)

[vas] is a command-line tool for [V.A.S].
It gets your installed packages informations (its names & versions) and transport to [V.A.S] and vulnerability diagnostic.

## Get a token key

1. visit https://vas.lepra.jp
2. Sign up or Login
3. `Account` click
4. You can find a token key!!

## Command Line Interface

### Download and Build

```
$ go get github.com/VulnerabilityAlertService/vas@latest
```

### Binaries

See [latest release](https://github.com/VulnerabilityAlertService/vas/releases/latest).

### Usage

```
$ vas -h
Transport installed packages info(name & version) to V.A.S.(https://vas.lepra.jp)

Usage:
  vas [flags]
  vas [command]

Available Commands:
  config      Edit config
  help        Help about any command
  version     Print version

Flags:
  -h, --help           help for vas
  -t, --token string   set a token key

Use "vas [command] --help" for more information about a command.
```

#### Save token key and Execute

```
$ vas config
Enter your token key: your-token-key
Update successfully!

$ vas
====== VulnerabilityAlertService ======
------        Success!!          ------
Check it out! https://vas.lepra.jp
=======================================
```

#### NOT Save token key and Execute

```
$ vas -t your-token-key
====== VulnerabilityAlertService ======
------        Success!!          ------
Check it out! https://vas.lepra.jp
=======================================
```

#### Use in `docker build`

##### Add config to your Dockerfile

```dockerfile
ARG TOKEN
RUN curl -s https://raw.githubusercontent.com/VulnerabilityAlertService/vas/master/executer.sh | sh
```

##### Build

```sh
$ docker build --build-arg TOKEN=yourtoken -f your_Dockerfile .
```

[vas]: https://github.com/VulnerabilityAlertService/vas "VulnerabilityAlertService/vas: vas - Vulnerability diagnostic tool"
[V.A.S]: https://vas.lepra.jp "Vulnerability Alert Service"