# Building Context Functionalities

## 1. `context.Background`
This is used to create the root of a context tree. It is never (or can not be) cancelled, has no values, and has no deadline
usually initialized in main() and passed down topmost function calls.

## 2. `context.WithCancel`
This is used to create a context that can be cancelled. It returns a context and a cancel function. The cancel function should be called when the context is no longer needed.
