# Pair of Generators 
This program calculates solves the challenge described in [this link](https://docs.google.com/document/d/1eurMI2z6JIfHM9zPg4gWz5NzB8TQxOmSAE4GN_3C8co/edit?usp=sharing).

## Requirements 

To run this program, clone the repository into your local environment. 

```sh
$ git clone https://github.com/ksetoue/pair-generators.git
```

Before you run, you'll need to have Go installed. [Follow this guide to install and configure your environment](https://golang.org/doc/install). 

Inside the folder you cloned from github, run the following on your terminal to execute the application:  
```sh
$ go run main.go
```

To execute tests, run:
```sh
$ go test -v 
```

This application uses the example described on the document mentioned above, in which: 

```
A start value is 65 
B start value is 8921

A factor is 16807
B factor is 48271

product is devided by 2147483647
```


