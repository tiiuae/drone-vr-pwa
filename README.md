# Drone VR PWA

This is the drone VR progressive web application, which works with
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

## VR PWA Setup

On the Pixel phone, visit the server's IP address on the Chrome browser. It
should load a black page with a stream from Gazebo.

Next, click on the menu on the top right corner of the browser, and select "Add
to Home screen."

The progressive web application should appear on the home screen.

**Note:** Sometimes the image may be frozen or blank, you may swipe down to
refresh the page.

You also need to set screen timeout in the phone's settings to the highest option.

If the machine IP address changes, you will have to remove the app from the home
screen and add it again.

