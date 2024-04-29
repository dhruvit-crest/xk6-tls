# xk6-tls

A k6 extension for sending data to tcp tls port

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:

  ```shell
  go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:

  ```shell
  xk6 build master \
    --with github.com/dhruvit96/xk6-tls
  ```

## Example

```javascript
import tls from 'k6/x/tls';
import { check } from 'k6';

const conn = tls.connect('host:port');

export default function () {
  tls.write(conn, 'Say Hello');
  let res = String.fromCharCode(...tls.read(conn, 1024))
  check (res, {
    'verify ag tag': (res) => res.includes('Hello')
  });
}
```
