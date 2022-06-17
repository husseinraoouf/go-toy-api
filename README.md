<h1 align="center">Scenario</h1>

<p align="center">
  <img src="https://toggl.notion.site/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F8bbb995e-e2fb-48aa-91e3-880cb5f5b12c%2FUntitled.png?table=block&id=296fa052-31c7-4709-9ac9-f41b501046c8&spaceId=16d29cae-5260-486d-ac3b-559da6a43a25&width=380&userId=&cache=v2" />
</p>

## How to setup

install dependencies and tools

```shell
make install_tools
```

## How to build

it will run on port **8080**

```shell
make build
```

## How to start

To start the app you must prepare database first

1. Prepare Database First

   ```shell
   make database_migrate
   ```

2. Run the api

   it will run on port **8080**

   ```shell
   make run
   ```

## How to test

```shell
make test
```

### To see coverage

```shell
make test_coverage
```

## Other Commands Commands

### To See All Avialable Commands

```shell
make help
```
