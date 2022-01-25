# hashdir

A simple tool to calculate the hash of a directory. Only considering it's
contents, ignoring all attributes.

Basically a copy of <https://github.com/gosimple/hashdir> with some
improvements, bug fixes & a CLI tool to wrap the library.

## CLI Usage

Install the cli tool by downloading the appropriate binary / archive from
<https://github.com/brad-jones/hashdir/releases>

Alternative installation methods:

- **[HomeBrew](https://brew.sh/):** `brew install brad-jones/tap/hashdir`
- **[Scoop](https://scoop.sh/):**
  `scoop bucket add brad-jones https://github.com/brad-jones/scoop-bucket.git; scoop install hashdir;`
- **[Docker](https://www.docker.com/):**
  `docker pull ghcr.io/brad-jones/hashdir/cli:latest`

Then just execute `hashdir` with a path and an optional hash algorithm
_(defaults to `sha256`)_.

```
hashdir [-alg md5|sha1|sha256|sha512] /a/path
```

The hash will be output as a hex encoded string, with no other output, not even
a newline.

> TIP: A docker run command might look like:
> `docker run --rm -v /a/path:/a/path ghcr.io/brad-jones/hashdir/cli:latest /a/path`

## Library Usage

Consume the go module like any other.

```go
package main

import (
  "github.com/brad-jones/hashdir"
)

func main() {
  hash, err := hashdir.Make("/a/path", "md5|sha1|sha256|sha512")
  if err != nil {
    panic(err)
  }

  // do something with "hash"...
}
```
