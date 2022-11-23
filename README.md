# (Go)SheepIt Render Farm Client

![build pipeline](https://github.com/stensonb/go-sheepit-client/actions/workflows/main.yml/badge.svg)

## Overview

This is a fork of the go-sheepit-client written found [here](https://github.com/stuarta0/go-sheepit-client).

**NOTE: This project is a work in progress and will eventually be used as a replacement for the official Java client**

The purpose of this client is to provide a native executable for each platform that interacts with the distributed render farm [SheepIt](https://www.sheepit-renderfarm.com/). This removes the dependency on the JVM which reduces the required overhead and makes server deployment easier.

The goals for this project include:
* Feature-parity with the official Java client
* Complete test coverage
* Stability
* Extensibility

## Compilation

    go get github.com/stensonb/go-sheepit-client/...

## Usage

    gosheepit.exe -help

When you are doing development work, you can use a mirror of the main site specifically made for demo/dev. The mirror is located at **http://sandbox.sheepit-renderfarm.com**, and you can use it by passing `-server http://sandbox.sheepit-renderfarm.com` to your invocation of the client.

## Contributing
Contributions are welcome.  Loose requirements for contributions include:
* unit tests covering your proposed changes
* excellent documentation describing your proposed changes

Please file an issue, and submit a PR against this repo for review.

## Current Roadmap
### v1.0.0
* successfully rendering scenes on linux machines via cpu
* released pre-built binaries to make installation/execution easy

## Future Roadmap Items
* successfully rendering scenes on non-linux machines via cpu
* successfully rendering scenes on machines via gpu
* UI for gosheepit delivered over http server on localhost
* pluggable storage backend? s3? ipfs?
* deployed as docker container?
* helm chart? 
* teach client about non-sheepit servers?
