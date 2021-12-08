# go_websocket
https://github.com/tsawler/ws-udemy/releases/tag/v12
## notes

### V1
使用websocket的線上聊天，列出使用者及其訊息，斷線重連。
![Imgur](https://i.imgur.com/qPGDjok.jpg)

## go libs
go get github.com/CloudyKit/jet/v6

go get github.com/bmizerany/pat

go get github.com/gorilla/websocket

https://github.com/CloudyKit/jet/v6

https://github.com/bmizerany/pat


### javascript libs

* [notie](https://github.com/jaredreich/notie)
  ：讓alert message更好看

* [reconnecting-websocket](https://github.com/joewalnes/reconnecting-websocket)

若server斷線，js會重新連線：

```js
var ws = new WebSocket('ws://....');
// you can replace with:
var ws = new ReconnectingWebSocket('ws://....');
```

## REF
* https://www.udemy.com/course/working-with-websockets-in-go/

## plugin
EmmetEverywhere Plugin for IntelliJ IDEA
