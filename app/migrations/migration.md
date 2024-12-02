# Migration

## Tools utilized:

- Makefile

- github.com/pressly/goose



## Create

In the root directory, run:

```shell
make migrate input your migration name.
```

Then, input the migration name.



## Up

In the root directory, run:

```shell
make up
```

There's other options for migration up:

- up by one version

- up to specific version

### Up by One

Run:

```shell
make up-one
```

### Up to

Run:

```shell
make up-to
```

Then, input the version desired.



## Down

In the root directory, run:

```shell
make down
```

There's other option for migration down:

- down to specific version

### Down to

Run:

```shell
make down-to
```

Then, input the version desired.



## Rollback

In the root directory, run:

```shell
make reset
```



## Re-run latest migration

In the root directory, run:

```shell
make redo
```



## Status

In the root directory, run:

```shell
make status
```



## Version

In the root directory, run:

```shell
make version
```
