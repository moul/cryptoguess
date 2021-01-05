# cryptoguess

:smile: cryptoguess automatically detects and parses cryptography keys from files

[![CircleCI](https://circleci.com/gh/moul/cryptoguess.svg?style=shield)](https://circleci.com/gh/moul/cryptoguess)
[![GoDoc](https://godoc.org/moul.io/cryptoguess?status.svg)](https://godoc.org/moul.io/cryptoguess)
[![License](https://img.shields.io/github/license/moul/cryptoguess.svg)](https://github.com/moul/cryptoguess/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/moul/cryptoguess.svg)](https://github.com/moul/cryptoguess/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/cryptoguess)](https://goreportcard.com/report/moul.io/cryptoguess)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/cryptoguess/badge)](https://www.codefactor.io/repository/github/moul/cryptoguess)
[![codecov](https://codecov.io/gh/moul/cryptoguess/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/cryptoguess)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/cryptoguess.svg)](https://microbadger.com/images/moul/cryptoguess)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)


## Usage

```console
$ find test/ -type f | xargs cryptoguess
test/pem-rsa-pubkey.txt:            potential candidates: PEM encoded data: x509: DER encoded public key, PEM encoded data
test/jwt-token.txt:                 JWT signed token
test/ssh-rsa-authorized-key.txt:    SSH authorized key
test/rsa-pubkey.txt:                potential candidates: BASE64 encoded data: x509: DER encoded public key, BASE64 encoded data
test/crypto-memory/D.der:           x509: PKCS#1 public key (RSA) in ASN.1 DER form
test/crypto-memory/E:               PEM encoded data
test/crypto-memory/A.pub:           SSH authorized key
test/crypto-memory/A:               PEM encoded data
test/crypto-memory/B.pem:           potential candidates: PEM encoded data: x509: PKCS#1 public key (RSA) in ASN.1 DER form, PEM encoded data
test/crypto-memory/B.pub:           SSH authorized key
test/crypto-memory/D.with-password: PEM encoded data
test/crypto-memory/C.pub:           SSH authorized key
test/crypto-memory/D:               potential candidates: PEM encoded data: x509: PKCS#1 private key (RSA) in ASN.1 DER form, PEM encoded data
test/crypto-memory/D.pub:           SSH authorized key
test/crypto-memory/A.der:           x509: PKCS#1 public key (RSA) in ASN.1 DER form
test/crypto-memory/B:               PEM encoded data
test/crypto-memory/C:               PEM encoded data
test/crypto-memory/B.der:           x509: PKCS#1 public key (RSA) in ASN.1 DER form
test/crypto-memory/F.pem:           potential candidates: PEM encoded data: x509: PKCS#1 public key (RSA) in ASN.1 DER form, PEM encoded data
test/crypto-memory/D.pem:           potential candidates: PEM encoded data: x509: PKCS#1 public key (RSA) in ASN.1 DER form, PEM encoded data
test/crypto-memory/F.pub:           SSH authorized key
test/crypto-memory/A.pem:           potential candidates: PEM encoded data: x509: PKCS#1 public key (RSA) in ASN.1 DER form, PEM encoded data
test/crypto-memory/F.der:           x509: PKCS#1 public key (RSA) in ASN.1 DER form
```

```console
$ find test/ -type f | xargs file
test/pem-rsa-pubkey.txt:            ASCII text
test/jwt-token.txt:                 ASCII text, with very long lines, with no line terminators
test/ssh-rsa-authorized-key.txt:    OpenSSH RSA public key
test/rsa-pubkey.txt:                ASCII text, with very long lines, with no line terminators
test/crypto-memory/D.der:           data
test/crypto-memory/E:               OpenSSH private key
test/crypto-memory/A.pub:           OpenSSH RSA public key
test/crypto-memory/A:               OpenSSH private key
test/crypto-memory/B.pem:           ASCII text
test/crypto-memory/B.pub:           OpenSSH RSA public key
test/crypto-memory/D.with-password: PEM RSA private key
test/crypto-memory/C.pub:           OpenSSH ED25519 public key
test/crypto-memory/D:               PEM RSA private key
test/crypto-memory/D.pub:           OpenSSH RSA public key
test/crypto-memory/A.der:           data
test/crypto-memory/B:               OpenSSH private key
test/crypto-memory/C:               OpenSSH private key
test/crypto-memory/B.der:           data
test/crypto-memory/F.pem:           ASCII text
test/crypto-memory/D.pem:           ASCII text
test/crypto-memory/F.pub:           OpenSSH RSA public key
test/crypto-memory/A.pem:           ASCII text
test/crypto-memory/F.der:           data
```

---

```console
$ cryptoguess --debug test/ssh-rsa-authorized-key.txt
test/ssh-rsa-authorized-key.txt: SSH authorized key
- PEM encoded data: err: no PEM data found
- SSH authorized key: *cryptoguess.ParsedSSHAuthorizedKey: &{0xc00005c8c0 lorem ipsum []}
- x509 DER encoded public key: err: asn1: structure error: tags don't match (16 vs {class:1 tag:19 length:115 isCompound:true}) {optional:false explicit:false application:false private:false defaultValue:<nil> tag:<nil> stringType:0 timeType:0 set:false omitEmpty:false} publicKeyInfo @2
```

---

```console
$ cryptoguess -h
NAME:
   cryptoguess - A new cli application

USAGE:
   cryptoguess [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -D    (default: false)
   --list, -l     (default: false)
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Decoders

| Encoding                         | Status             | Recursive          |
|----------------------------------|--------------------|--------------------|
| aes                              | :red_circle:       | :red_circle:       |
| ascii85                          | :red_circle:       | :red_circle:       |
| asn1                             | :red_circle:       | :red_circle:       |
| base32                           | :red_circle:       | :red_circle:       |
| base64                           | :white_check_mark: | :white_check_mark: |
| cipher                           | :red_circle:       | :red_circle:       |
| csv                              | :red_circle:       | :red_circle:       |
| des                              | :red_circle:       | :red_circle:       |
| dsa                              | :red_circle:       | :red_circle:       |
| ecdsa                            | :red_circle:       | :red_circle:       |
| elliptic                         | :red_circle:       | :red_circle:       |
| encodings (utf-8)                | :red_circle:       | :red_circle:       |
| encrypted jwt                    | :red_circle:       | :red_circle:       |
| gob                              | :red_circle:       | :red_circle:       |
| gzip,lzw,...                     | :red_circle:       | :red_circle:       |
| json                             | :red_circle:       | :red_circle:       |
| pem                              | :white_check_mark: | :white_check_mark: |
| rsa                              | :red_circle:       | :red_circle:       |
| signed jwt                       | :white_check_mark: | :red_circle:       |
| ssh                              | :white_check_mark: | :red_circle:       |
| tls                              | :red_circle:       | :red_circle:       |
| url escaped                      | :red_circle:       | :red_circle:       |
| x509: DER certificate list       | :white_check_mark: | n/a                |
| x509: Elliptic Curve private key | :white_check_mark: | n/a                |
| x509: PKCS#1 RSA private key     | :white_check_mark: | n/a                |
| x509: PKCS#8 private key         | :white_check_mark: | n/a                |
| x509: PKCS#8 public key          | :white_check_mark: | n/a                |
| x509: PKIX public key            | :white_check_mark: | n/a                |
| x509: certificate                | :white_check_mark: | n/a                |
| x509: certificate list           | :white_check_mark: | n/a                |
| x509: certificate request        | :white_check_mark: | n/a                |
| x509: certificates               | :white_check_mark: | n/a                |
| xml                              | :red_circle:       | :red_circle:       |


## Install

### CLI

```console
$ go get -u moul.io/cryptoguess
```

### Library

```console
$ go get -u moul.io/cryptoguess/cryptoguess
```

## As a library

See https://godoc.org/moul.io/cryptoguess/cryptoguess

## License

© 2019-2021 [Manfred Touron](https://manfred.life) -
[Apache-2.0 License](https://github.com/moul/cryptoguess/blob/master/LICENSE)
