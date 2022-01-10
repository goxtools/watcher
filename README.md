# watcher any for Golang

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
-        defer gorecover.GoRecover(context.Background(), c.GetConsumerName()+"job异常")
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
    go watch.Watcher(func() {
        ctx := new(Context)
        c.Consumer(ctx)
    })
}
```

# Test Case
```go

func TestSen(t *testing.T) {

	ti := time.NewTicker(time.Second * 1)
	go Watcher(func(args ...interface{}) {
		for {
			select {
			case <-ti.C:
				fmt.Println(time.Now().Second())
				if time.Now().Second()%5 == 0 {
					panic("123312")
				}
			default:
				time.Sleep(time.Millisecond * 50)
			}
		}
	})

	for {
		select {

		}
	}
}
```