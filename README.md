# watcher any for Golang

[![Go Tests](https://github.com/goxtools/watcher/actions/workflows/go.test.yml/badge.svg)](https://github.com/goxtools/watcher/actions/workflows/go.test.yml)

# Installation

```go
go get github.com/goxtools/watcher@v0.0.1
```

# Quickstart

### for consumer [old version]
```diff
- ch := make(chan int, c.GetConsumerNumber())
- for {
-    ch <- 1
-    go func(channel chan int, c mq.Consumer) {
-        defer gorecover.GoRecover(context.Background(), c.GetConsumerName()+"job err")
-        ctx := new(Context)
-        if err := c.Consumer(ctx); err != nil {
-            <-channel
-            return
-        }
-        <-channel
-    }(ch, c)
-}
```

### for consumer [watcher]
```go
for i := 0; i < c.GetConsumerNumber(); i++{
    go func() {
        w := NewWatcher(10, time.Second*1, time.Second*10)
        w.On(func(args ...interface{}) {
            c.Consumer(new(Context))
        }
    }
}
```

# Test Case
```go
func TestWatcher(t *testing.T) {
    rand.Seed(time.Now().UnixNano())

    //	Retry 10 times, the retry interval is 1 second,
    //	after 10 seconds of stable operation,
    //	reset the number of retries
    w := NewWatcher(10, time.Second*1, time.Second*10)
    w.On(func(args ...interface{}) {
        for {
            if rand.Int()%9 == 0 {
                panic("crash")
            }
            time.Sleep(500 * time.Millisecond)
            fmt.Printf("%#v\n", w)
        }
    })
}
```
