# notfound

This is web application what listen  specified port and always return only response `404 not found`.

## Motivation

When using a docker-compose project to develop a project that contains multiple web applications, use it as a dummy application that will change the application that is not used.

## Install

```
$ go install github.com/ken39arg/notfound
```

### Docker image

```
$ docker pull ken39arg/notfound
$ docker run -p 8080:8080 ken39arg/notfound -port=8080
```

## Usage

```
$ notfound -port=8080
# or
$ PORT=8080 notfound
```

## Licence

[MIT](./LICENSE)

## Author

[ken39arg](https://github.com/ken39arg)
