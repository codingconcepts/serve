# serve
An almost offensively simple file server.

### Installation

Find the release that matches your architecture on the [releases](https://github.com/codingconcepts/serve/releases) page.

Download the tar, extract the executable, and move it into your PATH:

```
$ tar -xvf serve_[VERSION]_macos.tar.gz
```

### Usage
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