# Password protection for static sites using Go and WebAssembly

The concept of storing a static site under a directory with a name of a hashed password is borrowed from [matteobrusa](https://github.com/matteobrusa/Password-protection-for-static-pages). 

The repo is modified to use Go and WebAssembly for calculating a hash. The auth UI and JavaScript are simplified. No need to fetch external resources except provided with this repository.

## File structure

```
18a43b7c6d7cd4316a3320229ccb362ff56ffc8712db9dd65b0c01a5f744ed4d    <- hash of a password
index.html                                                          <- auth UI
hash.wasm                                                           <- go module
wasm_exec.js                                                        <- go wasm js glue
```

## Demo

For a demo: 

- download the repository, 
- go inside the `./demo`,
- run `goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'`,
- check the site at `localhost:8080`.

Run `go get -u github.com/shurcooL/goexec` to get `goexec` if you don't have one.

## Deployment

For use you don't need to install Go or anything. Just do the following:

- copy the content of `./demo` into your folder on a server,
- and put your site under the name of a hashed with SHA512 password.

In order to get the hash of your password, use https://play.golang.org/p/0i2VETHvxJL — change "your password" with your text.

The code behind the link is simple:

``` go
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := sha512.New512_256()
	_, err := io.WriteString(h, "your password")
	if err != nil {
		fmt.Println(err)
	}

	calculated := hex.EncodeToString(h.Sum(nil))
	fmt.Println(calculated)
}
```
