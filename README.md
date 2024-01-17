# ENULS Network
ENULS, initiated by the NULS community, is a blockchain that is fully compatible with EVM and Web3 API interfaces. ENULS uses the NULS as the main asset in its network. The contract assets and chain assets from NULS can also be bridge to the ENULS network.

ENULS can be also seen as a layer2 scaling solution for Ethereum that brings better performance and lower gas-fee cost. Meanwhile, ENULS blockchain will also be supporting efficient deployment of the Ethereum ecosystem projects.


Go ENULS project is developed based on Go-Ethereum.
 
## Building The Source

For prerequisites and detailed build instructions please read the [Installation Instructions](https://geth.ethereum.org/docs/install-and-build/installing-geth).

Building `geth` requires both a Go (version 1.16 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The go-ethereum project comes with several wrappers/executables found in the `cmd`
directory.

|    Command    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| :-----------: |----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  **`geth`**   | Our CLI client. It is the entry point into the ENULS network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the ENULS network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` and the [CLI page](https://geth.ethereum.org/docs/interface/command-line-options) for command line options. |

## Running `enuls`

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://geth.ethereum.org/docs/interface/command-line-options)),
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `geth` instance.

### Hardware Requirements

Minimum:

* CPU with 2+ cores
* 4GB RAM
* 1TB free storage space to sync the Mainnet
* 8 MBit/sec download Internet service

Recommended:

* Fast CPU with 4+ cores
* 16GB+ RAM
* High Performance SSD with at least 1TB free space
* 25+ MBit/sec download Internet service
 
### Full Node on The ENULS Network

```shell
$ geth --enuls console
```

The `console` subcommand has the exact same meaning as above and they are equally
useful on the testnet too. Please, see above for their explanations if you've skipped here.
 
### Full Node on The ENULS Test Network
 

```shell
$ geth --enulstest console
```

## Contribute to NULS
We are committed to making blockchain technology simpler and our slogan is "NULS Making It Easier to Innovate".

Get to know NULS developers
https://nuls.io/developer

You are welcome to contribute to NULS! We sincerely invite developers with rich experience in the blockchain field to join the NULS technology community.
https://nuls.io/community

Documentation：https://docs.nuls.io

NULS Brand Assets: https://nuls.io/brand-assets



## License

NULS is released under the [MIT](http://opensource.org/licenses/MIT) license.
Modules added in the future may be release under different license, will specified in the module library path.

## Community

- Website: https://nuls.io
- Twitter: https://twitter.com/nuls
- Discord:https://discord.gg/aRCwbj47WN
- Telegram: https://t.me/Nulsio)
- Medium: https://nuls.medium.com
- Forum: https://forum.nuls.io
- GitHub: https://github.com/nuls-io

####  
