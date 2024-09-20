# upsPowerShow

A tray program that shows Greengroup's UPS battery information.

一个可以显示绿联UPS电池信息的托盘程序。

**Before you use, please read the following introduction.**

You should change the `userId` and `upsId` in `websocketclient/wss.go`.

Please be advised, you may still need to capture the packet of the mobile APP. Because I don't have any other ways to fetch those two Ids above.

**在你使用前，请务必阅读以下说明**

你需要在 `websocketclient/wss.go` 文件下修改 `userId` 和 `upsId` 方可使用。

具体操作是抓手机APP“绿联储能”的包，具体操作恕不赘述。

## Available OS

It is designed to be working on Windows. I don't know if it can run normally on other OS.

## 可用平台

本软件是为Windows设计的，别的操作系统能否运行我不清楚。

## Compile by yourself

```
go build main.go
```

You can add some parameters if you prefer. e.g. To disable console on Windows, run this with `-ldflags="-H windowsgui"`.

## 自行编译

```
go build main.go
```

你可以自行添加参数，如Windows上可添加 `-ldflags="-H windowsgui"` 以关闭控制台窗口。

## Download from Release

Or, you can directly download from release page which is much easier and more convenient.

## 直接下载

或者，你可以直接在Release页下载，更方便、更简单。

## How to use

Just click the .exe, set the tray to always show on the screen, there you go!

If you like, you can copy a shortcut to your Startup directory so that it will start automatically when booting.

But ensure the `icons/` needs to be in the same directory with application.

## 如何使用

双击exe文件即可，再在任务栏设置中将其设定为始终显示。

你也可以把快捷方式放到开机自启的文件夹内。

但确保 `icons/` 文件夹与主程序在同一目录下。