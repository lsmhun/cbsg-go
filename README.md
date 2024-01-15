# Corporate Bullshit Generator for Go
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/lsmhun/cbsg-go/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/lsmhun/cbsg-go/tree/main)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flsmhun%2Fcbsg-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Flsmhun%2Fcbsg-go?ref=badge_shield)


This is the [Corporate Bullshit Generator](http://cbsg.sf.net/) rewritten for Golang.
Implementation is based on [Corporate Bullshit Generator for Java](https://github.com/lsmhun/cbsg-java).
Transcoding committed by ChatGPT 3.5

* [English documentation](./docs/descr_en.md)
* [Hungarian documentation](./docs/descr_hu.md)

# Features

This can be used as a standalone program or a library. You can generate other dictionaries based
on the [cbsg dictionary file](./dict/en/cbsg_dictionary.csv) .

# Usage

## Build application
You can build as a simple standalone application with built-in dictionary.
Default value is "shortWorkshop".
```shell
$ make build
$ ./out/bin/cbsg --help
Available options:
--workshop
--shortWorkshop
--financialReport
--sentenceGuaranteedAmount=<ANY_INTEGER>

--dictionaryFile=<DICTIONARY_FILE>
--help
```
```
$ ./out/bin/cbsg
Controlling should be committed across industry sectors. A well-planned and fast-evolving dematerialization carefully promotes the decision makers from the get-go. The Customer Experience Management efficiently cost-control measures ensuring market opportunities. Our wide-range market conditions diligently enables the white-collar workers in the core. The adjacencies an evolutionary and executive-level silo resulting in a long-term run-rate efficiency.

```
## Using as a library
You can find a corporate [example](./examples/basic/main.go) with Dilbert's pointy haired boss.

Example output with [cowsay](https://github.com/Code-Hex/Neo-cowsay) (with added ASCII art):
```
 __________________________________________ 
/ There can be no gain in task efficiency  \
| until we can achieve a rapid growth      |
| momentum. Decision-maker and             |
| relationship result ins                  |
| innovation-driven and cross-enterprise   |
| style guidelines across the board. Our   |
| marketplace and right emotional impact   |
| organically turbocharges a matrix across |
| the organizations. The organization our  |
| key performance indicators on a          |
| transitional basis. The corporate values |
| cost-effectively strengthen the thought  |
\ leaders across the board.                /
 ------------------------------------------
          \
           \
              @         @
             @@  ..-..  @@
             @@@' _ _ '@@@
              @(  oo   )@
               |  (_)  |
               |   _   |
               |_     _|
              /|_'---'_|\
             / | '\_/' | \
            /  |  | |  |  \
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flsmhun%2Fcbsg-go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Flsmhun%2Fcbsg-go?ref=badge_large)
