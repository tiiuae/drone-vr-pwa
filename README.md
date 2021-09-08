# Drone VR PWA

This is the drone VR PWA application, which works with
[jgoppert/purt\_catkin\_ws].

## Prerequisites

You need the following:

- [jgoppert/purt\_catkin\_ws]
- [Go]

[jgoppert/purt\_catkin\_ws]: https://github.com/jgoppert/purt_catkin_ws
[Go]: https://golang.org

## Usage

First, you will need to launch the Abu Dhabi world.

You may compile and run with one command:

```sh
go run main.go
```

Port may be optionally defined, and IP address is automatically
determined. You may override them:

```sh
go run main.go 8181 10.0.0.105
```


