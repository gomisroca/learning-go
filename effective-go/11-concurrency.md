## Concurrency

### Share by Communicating

Go encourages sharing values on channels, and never actively shared by separate threads of execution. Only one goroutine has access to the value at any given time. Data races cannot occur.
