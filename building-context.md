# Building Context Functionalities

## 1. `context.Background`
This is used to create the root of a context tree. It is never (or can not be) cancelled, has no values, and has no deadline
usually initialized in main() and passed down topmost function calls.

## 2. `context.WithCancel`
This is used to create a context that can be cancelled. It returns a context and a cancel function. The cancel function should be called when the context is no longer needed.
When we cancel a context, all contexts derived from it are also cancelled.
Cancelling a child context does not cancel the parent context.

## 3. `context.WithDeadline`
This is used to create a context that is cancelled when the deadline is reached. It returns a context and a cancel function. The cancel function should be called when the context is no longer needed.
When the deadline is reached, the context is cancelled and all contexts derived from it are also cancelled.