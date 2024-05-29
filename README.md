kobrad
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/Pyrinpyi/pyipad)

kobrad is the reference full node kobra implementation written in Go (golang).
(Fork of https://github.com/kaspanet/kaspad)

## What is kobra

kobra is an attempt at a proof-of-work cryptocurrency with instant confirmations and sub-second block times. It is based on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a generalization of Nakamoto consensus.

## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install pyipad including all dependencies:

```bash
$ git clone https://github.com/kobradag/kobrad
$ cd kobrad
$ go install . ./cmd/...
```

- pyipad (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.


## Getting Started

kobrad has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ kobrad
```

## Discord
Join our discord server using the following link : https://discord.gg/ahnm5zENqg

## Issue Tracker

The [integrated github issue tracker](https://github.com/kobradag/kobrad/issues)
is used for this project.

## Documentation

The [documentation](https://github.com/kobradag/docs) is a work-in-progress

## License

pyipad is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
# kobrad
 
