# (Go)SheepIt Render Farm Client

## Overview

This is a fork of the go-sheepit-client written found [here](https://github.com/stuarta0/go-sheepit-client).

**NOTE: This project is a work in progress and will eventualy be used as a replacement for the official Java client**

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
* successfully rendering scenes on linux machines

## Future Roadmap Items
* UI for gosheepit delivered over http server on localhost
