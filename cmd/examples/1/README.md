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

### 2. Deadline
- Go was also supposed to work as a language for microservices or systems interacting with external resources.
- Whenever we are interacting with external systems, it is better to set a deadline for that interaction 


### 3. Err 
- Err returns one of two types of values. Either `context.Canceled` error or `context.DeadlineExceeded`.
- It just differenciates why the context says that work needs to be stopped/cancelled

### 4. Value
- Probably the only field that makes sense of the package name
- Extensively used in server side code to store request scoped variables like "authenticated_user","trace_id", "role", etc.
