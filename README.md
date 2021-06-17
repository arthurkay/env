# GO env

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/arthurkay/env)
[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/arthurkay/env/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkay/env)](https://goreportcard.com/report/github.com/arthurkay/env)
[![Go](https://github.com/arthurkay/env/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/arthurkay/env/actions/workflows/go.yml)

Reading GO env files into the current process's environmental variables.

This is an importable package thats meant to be a light weight environmental variable reader from files.
The package allows you to set up your environment variables in a file that is kept outside your version control system.

The advantage of such an architecture is so you can keep all your secrets "secret"!.

# How To Use env

You only need to import and load the module in your application to use it. e.g.


First create a `.env` file.

```.env
TOKEN=SuperSecretToken
```

If the file is in the root directory of your module, it will automatically be picked no need to pass it as na argument to the `Load()` function.

```go
package main

import (
    "github.com/arthurkay/env"
    "fmt"
    "os"
)

func main() {
    env.Load()
    fmt.Printf("Application token key is %s", os.Getenv("TOKEN"))
}

```

## NOTE:

If you name your `.env` file something else, please make sure you add  it in the load function as a string argument.

```go

...
env.Load("env_file")
...

```

You can even pass multiple files to the Load function. e.g

```go

...
env.Load("env", "env_file")
...


```


Dont forget to add the `.env` file to your `.gitignore` file so the `.env` file is not included in your commits.
