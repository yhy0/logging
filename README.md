## logging

自用 log 模块

```go


// screenReport 屏幕是否输出路径信息
logging.New(true, "", "Jie", false)
logging.Logger.Infoln("test")
logging.Logger.Debugln("debug")

```

在当前目录创建一个 logs 目录，记录日志信息

![image-20230316132709613](images/image-20230316132709613.png)
