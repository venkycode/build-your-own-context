# Building Context Functionalities

## 1. `context.Background`
This is used to create the root of a context tree. It is never (or can not be) cancelled, has no values, and has no deadline
usually initialized in main() and passed down topmost function calls.
