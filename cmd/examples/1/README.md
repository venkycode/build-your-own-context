# Need for context 

## The context.Context interface 

```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}   
```

## Understanding each Method

### 1. Done
- Go's focus on concurrency resulting in need of interrupting spawned goroutines 
- Done's return value tells if the parent function in the chain has asked to stop the work 

