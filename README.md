`waiter`
--------

Waits until the lock-file is removed. When starting the application a lock-file (`--lock-file-path`)
is created, and when the file is removed the `waiter` stops gracefully. When timeout is reached, the
application exists on error.

## Usage

```bash
waiter start
```

And:

```bash
waiter done
```