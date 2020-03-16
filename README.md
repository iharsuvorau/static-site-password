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
