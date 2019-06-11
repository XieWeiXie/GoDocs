# Time

- 延时系统

```
chan<- 接收
<-chan 发送


var rateLimiter = time.Tick(1 * time.Second)
<- rateLimiter


或者：

<- time.After(time.Second)
```