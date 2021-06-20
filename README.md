`waiter`
--------

In a nutshell, it waits until the lock-file is removed. When starting the application a lock-file
(`--lock-file-path`) is created, and when the file is removed the `waiter` stops gracefully. When
timeout is reached, the application exists on error.

## Usage

Please consider `--help` to see the possible flags, the possible sub-commands are:

```sh
waiter start
```

And:

```sh
waiter done
```

## Contributing

This project uses [GNU/Make][make] to concentrate all basic actions against the code-base. In order
to build the application, run:

```sh
make
```

To execute the project tests, run:

```sh
make test
```

And building a container-image out of local changes can be achieved by:

```sh
make image
```

Also consider the top-variables on the `Makefile` to tweak execution parameters.

[make]: https://www.gnu.org/software/make/