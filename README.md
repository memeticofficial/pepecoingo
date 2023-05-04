<div align="center">
  <img src="resources/pepecoinlogo.png?raw=true">
</div>

---

Node implementation for the pepecoin network.
a blockchains platform with high throughput, and blazing fast transactions.

## Installation

Pepecoin is an incredibly lightweight protocol, so the minimum computer requirements are quite modest.
Note that as network usage increases, hardware requirements may change.

The minimum recommended hardware specification for nodes connected to Mainnet is:

- CPU: Equivalent of 8 AWS vCPU
- RAM: 16 GiB
- Storage: 1 TiB
- OS: Ubuntu 20.04/22.04 or macOS >= 12
- Network: Reliable IPv4 or IPv6 network connection, with an open public port.

If you plan to build PepecoinGo from source, you will also need the following software:

- [Go](https://golang.org/doc/install) version >= 1.19.6
- [gcc](https://gcc.gnu.org/)
- g++

### Building From Source

#### Clone The Repository

Clone the PepecoinGo repository:

```sh
git clone git@github.com:memeticofficial/pepecoingo.git
cd pepecoingo
```

This will clone and checkout the `master` branch.

#### Building PepecoinGo

Build PepecoinGo by running the build script:

```sh
./scripts/build.sh
```

The `pepecoingo` binary is now in the `build` directory. To run:

```sh
./build/pepecoingo
```


### Docker Install

Make sure Docker is installed on the machine - so commands like `docker run` etc. are available.

Building the Docker image of latest `pepecoingo` branch can be done by running:

```sh
./scripts/build_image.sh
```

To check the built image, run:

```sh
docker image ls
```

The image should be tagged as `memeticplatform/pepecoingo:xxxxxxxx`, where `xxxxxxxx` is the shortened commit of the Pepecoin source it was built from. To run the Pepecoin node, run:

```sh
docker run -ti -p 9650:9650 -p 9651:9651 memeticplatform/pepecoingo:xxxxxxxx /pepecoingo/build/pepecoingo
```

## Running Pepecoin

### Connecting to Mainnet

To connect to the Pepecoin Mainnet, run:

```sh
./build/pepecoingo
```

You should see some pretty ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

### Connecting to Toadville

To connect to the Toadville Testnet, run:

```sh
./build/pepecoingo --network-id=Toadville
```


## Bootstrapping

A node needs to catch up to the latest network state before it can participate in consensus and serve API calls. This process (called bootstrapping) currently takes several days for a new node connected to Mainnet.

Improvements that reduce the amount of time it takes to bootstrap are under development.

The bottleneck during bootstrapping is typically database IO. Using a more powerful CPU or increasing the database IOPS on the computer running a node will decrease the amount of time bootstrapping takes.

## Generating Code

PepecoinGo uses multiple tools to generate efficient and boilerplate code.

### Running protobuf codegen

To regenerate the protobuf go code, run `scripts/protobuf_codegen.sh` from the root of the repo.

This should only be necessary when upgrading protobuf versions or modifying .proto definition files.

To use this script, you must have [buf](https://docs.buf.build/installation) (v1.11.0), protoc-gen-go (v1.28.0) and protoc-gen-go-grpc (v1.2.0) installed.

To install the buf dependencies:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

If you have not already, you may need to add `$GOPATH/bin` to your `$PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

If you extract buf to ~/software/buf/bin, the following should work:

```sh
export PATH=$PATH:~/software/buf/bin/:~/go/bin
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go-grpc
scripts/protobuf_codegen.sh
```

For more information, refer to the [GRPC Golang Quick Start Guide](https://grpc.io/docs/languages/go/quickstart/).

### Running protobuf codegen from docker

```sh
docker build -t pepecoin:protobuf_codegen -f api/Dockerfile.buf .
docker run -t -i -v $(pwd):/opt/pepecoin -w/opt/pepecoin pepecoin:protobuf_codegen bash -c "scripts/protobuf_codegen.sh"
```

### Running mock codegen

To regenerate the [gomock](https://github.com/golang/mock) code, run `scripts/mock.gen.sh` from the root of the repo.

This should only be necessary when modifying exported interfaces or after modifying `scripts/mock.mockgen.txt`.

## Versioning

### Version Semantics

PepecoinGo is first and foremost a client for the Pepecoin network. The versioning of PepecoinGo follows that of the Pepecoin network.

- currently all versions in testing/testnet

### Library Compatibility Guarantees

Because PepecoinGo's version denotes the network version, it is expected that interfaces exported by PepecoinGo's packages may change in `Patch` version updates.

### API Compatibility Guarantees

APIs exposed when running PepecoinGo will maintain backwards compatibility, unless the functionality is explicitly deprecated and announced when removed.

## Supported Platforms

PepecoinGo can run on different platforms, with different support tiers:

- **Tier 1**: Fully supported by the maintainers, guaranteed to pass all tests including e2e and stress tests.
- **Tier 2**: Passes all unit and integration tests but not necessarily e2e tests.
- **Tier 3**: Builds but lightly tested (or not), considered _experimental_.
- **Not supported**: May not build and not tested, considered _unsafe_. To be supported in the future.

The following table lists currently supported platforms and their corresponding
PepecoinGo support tiers:

| Architecture | Operating system | Support tier  |
| :----------: | :--------------: | :-----------: |
|    amd64     |      Linux       |       1       |
|    arm64     |      Linux       |       2       |
|    amd64     |      Darwin      |       2       |
|    amd64     |     Windows      |       3       |
|     arm      |      Linux       | Not supported |
|     i386     |      Linux       | Not supported |
|    arm64     |      Darwin      | Not supported |

To officially support a new platform, one must satisfy the following requirements:

| PepecoinGo continuous integration | Tier 1  | Tier 2  | Tier 3  |
| ---------------------------------- | :-----: | :-----: | :-----: |
| Build passes                       | &check; | &check; | &check; |
| Unit and integration tests pass    | &check; | &check; |         |
| End-to-end and stress tests pass   | &check; |         |         |

## Security Bugs

**We and our community welcome responsible disclosures.**

Please refer to our [Security Policy](SECURITY.md) and [Security Advisories](https://github.com/memeticofficial/pepecoingo/security/advisories).
