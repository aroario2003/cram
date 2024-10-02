# CRAM Cli 

## Introduction 

TODO

# Security Notice

We are aware that when starting dbsock you must enter the password for the database in plain text and that could be a problem for a security critical environment.

## Multi-platform building

Go has a robust build system and therefore certain features are used to make this application multi-platform. To clarify, multi-platform means that this will not run natively on all platforms without certain modifications to the build system and codebase. You do not have to make these modifications as they are already made for you. However, it is worth knowing what they are.

## Go build tags and comment compiler directives

Build tags in go tell the compiler which version of your application you want to build, a go build tag takes the following form:

```
// +build linux
```

This tells the compiler that whenever it gets a tag linux, it should build files with that comment on top of them. You can also specify multiple tags in these comments.

There are also go compiler directives in comments, this tells the compiler which platform to compile for based on the current go runtime which boils down to the operating system. These take the form:

```
//go:build linux
```

This tells the compiler that it should build this file specifically when the runtime is linux, the same could be said for windows if it was windows in that comment.

### Code Duplication Side Effect

Because of the multi-platform nature of this application, their is quite a bit of code duplication in multiple functions throughout. However, this is absolutely necessary because functions will return different types depending on the operating system and therefore must be duplicated for the platforms supported.

## What is dbsock?

In order to increase efficiency we create a program that creates a unix doamin socket that maintains the database conection even while the cli isnt running. Dbsock is the program that creates that socket and connects it to the database.

## Building dbsock

### On Linux

In order to build dbsock on linux, from the root of the project enter these commands:

```shell
$ cd dbsock
$ go build -o dbsock
```

### On Windows

In order to build dbsock on windows, from the root of the project enter these commands:

```shell
$ cd dbsock
$ go build -o dbsock.exe
```

## Building the main program

### On Linux

From the root of the project do

```shell
$ go build -o main
```

### On Windows

From the root of the project do

```shell
$ go build -o main.exe
```

## Credits

TODO
