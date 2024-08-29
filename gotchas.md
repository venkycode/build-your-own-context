# Gotcha's regarding context 

- Any mutation to a (parent) context creates a new inherited context from the parent context.
    - [example](cmd/examples/2/main.go)
- Cancelling a context does not ensure that the fuctions receiving that context suddenly stop. It is totally upto the called functions in the chain or spawned goroutines to honor the context's status.
- context are "almost" unmutable
