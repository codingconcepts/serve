# serve
An almost offensively simple file server.

## Installation

```
$ go get -u github.com/codingconcepts/serve
```

## Usage
```
$ serve -h
Usage of serve:
  -d string
        relative or absolute path of directory to serve
  -p int

$ serve -d . -p 1234
Serving from: .
Listening on: 192.168.10.10:1234
```