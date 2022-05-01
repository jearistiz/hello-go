# On the GO - Learning

Initially I'm following the [tutorials](https://go.dev/learn/) published on GO's website.

I will place here any information I think its useful for me (or you) for later reference.

## Some useful `go` commands

```sh
$ go mod init [your-module-sourcecode-url]
```

Enables dependency tracking of your module by creating a `go.mod` file in the project's root directory.
This file contains your code dependencies as well as a path-URL to find your code [possibly online].

---

```sh
$ go mod tidy
```

From the official documendation:
*Tidy makes sure `go.mod` matches the source code in the module.
It adds any missing modules necessary to build the current module's
packages and dependencies, and it removes unused modules that
don't provide any relevant packages. It also adds any missing entries
to go.sum and removes any unnecessary ones*

---

```sh
go mod edit -replace example/greetings=../greetings
```

Tells go.mod where to find a local package. In this case it tells to find `example/greetings` in `../greetings`. This is only recommended for code that is used locally.
