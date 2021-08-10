# Httpixy


A fast tool to scan Host Headers Injection  written in Go

---

## Resources

- [Installation](#installation)
	- [from Source](#from-source)
	- [from GitHub](#from-github)
- [Usage](#usage)
	- [Basic Usage](#basic-usage)
	- [Flags](#flags)
	- [Target](#target)
		- [Single URL](#single-url)
		- [URLs from list](#urls-from-list)
		- [from Stdin](#from-stdin)
	- [Method](#method)
	- [Using Proxy](#using-proxy)
	- [Concurrency](#concurrency)
	- [Silent](#silent)
	- [Version](#version)
- [Help & Bugs](#help--bugs)
- [Version](#version)

## Installation

### from Source

If you have go1.13+ compiler installed and configured:

```bash
▶ GO111MODULE=on go get -v github.com/Zeckers/httpixy/cmd/httpixy
```

In order to update the tool, you can use `-u` flag with go get command.

### from GitHub

```bash
▶ git clone https://github.com/Zeckers/httpixy
▶ cd httpixy/cmd/httpixy
▶ go build .
▶ mv httpixy /usr/local/bin
```

## Usage

### Basic Usage

Simply, Httpixy can be run with:

```bash
▶ httpixy -u "http://target"
```

### Flags

```bash
▶ httpixy -h
```

This will display help for the tool. Here are all the switches it supports.

| Flag             	| Description                                    	|
|------------------	|------------------------------------------------	|
| -u, --url        	| Define single URL to test                      	|
| -l, --list       	| Test URLs within file                          	|
| -X, --method     	| Specify request method to use _(default: GET)_ 	|
| -o, --output     	| File to save results                              |
| -x, --proxy      	| Use specified proxy to fuzz                       |
| -c, --concurrent 	| Set the concurrency level _(default: 25)_      	|
| -s, --silent     	| Silent mode                                    	|
| -V, --version    	| Show current httpixy version                   	|
| -h, --help       	| Display its help                               	|

### Target

You can define a target in 3 ways:

#### Single URL

```bash
▶ httpixy -u "http://target"
```

#### URLs from list

```bash
▶ httpixy -l /path/to/urls.txt
```

#### from Stdin

In case you want to chained with other tools.

```bash
▶ subfinder -d target -silent | httpx -silent | httpixy
```

### Method

By default, httpixy makes requests with `GET` method.
If you want to change it, you can use the `-X` flag.

```bash
▶ httpixy -u "http://target" -X "GET"
```

### Output

You can also save fuzzing results to a file with `-o` flag.

```bash
▶ httpixy -l /path/to/urls.txt -o /path/to/results.txt
```

### Using Proxy

Using a proxy, proxy string can be specified with a `protocol://` prefix to specify alternative proxy protocols.

```bash
▶ httpixy -u "http://target" -x http://127.0.0.1:8080
```

### Concurrency

Concurrency is the number of fuzzing at the same time. Default value httpixy provide is `25`, you can change it by using `-c` flag.

```bash
▶ httpixy -l /path/to/urls.txt -c 50
```

### Silent

If you activate this silent mode with the `-s` flag, you will only see vulnerable targets.

```bash
▶ httpixy -l /path/to/urls.txt -s | tee vuln-urls.txt
```

### Version

To display the current version of httpixy with the `-V` flag.

```bash
▶ httpixy -V
```

## Help & Bugs

If you are still confused or found a bug, please [open the issue](https://github.com/Zeckers/httpixy/issues). All bug reports are appreciated, some features have not been tested yet due to lack of free time.


## Version

**Current version is 1.0.0** and still development.