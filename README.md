[![Go Reference](https://pkg.go.dev/badge/github.com/uudashr/pkgname.svg)](https://pkg.go.dev/github.com/uudashr/pkgname)

# pkgname
Linter and analyzer that detects package names that do not follow Go idiomatic naming conventions.

## Usage

Install the linter:
```shell
go install github.com/uudashr/pkgname/cmd/pkgnamecheck@latest
```

Run the linter
```shell
pkgnamecheck ./...
```

# References

> Good package names are short and clear. They are lower case, with no `under_scores` or `mixedCaps`. ...
>
> [The Go Blog: Package names](https://go.dev/blog/package-names#package-names)

> The style of names typical of another language might not be idiomatic in a Go program. Here are two examples of names that might be good style in other languages but do not fit well in Go:
> 
> - `computeServiceClient`
> - `priority_queue`
> 
> [The Go Blog: Package names](https://go.dev/blog/package-names#package-names)

> By convention, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps. ...
>
> [Effective Go - Package names](https://go.dev/doc/effective_go#package-names)

