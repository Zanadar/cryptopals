### Golang Crypto Solver

#### TODO:
- Modularize
- Build into CLI tool
  - Implement text file input and command line input in CLI
  - Implement text file and on screen output
  - Build some kind of status bar for CLI
- Build in AES-128 ecb with [this](https://github.com/golang/go/issues/5597)
##### What is does
This program solves breaks encrypted text. It's based on the exercises in the Matasano Cryptopals challenges. The idea is to build a CLI program that uses statistical analysys and tries to brute a decrypt using a few basic decrypt techniques, including XOR, single-byte XOR, repeating key XOR, and vulnerabilities in AES-128 ECB.

##### How to use it
Basically once it's installed, you should be able to run the command `crapto -f './mysecret.txt` and get some kind of decrypto, assuming its encrypted in something I know how to crack!
