# Queue

A FIFO queue backed by the standard library doubly linked list. Supports O(1) insertion and removal.

## Import

```go
import (
    "github.com/connorjacobsen/gorithms/containers/queue"
)
```

### API

```go
type Queue struct{}
```

Create a new Queue:

```go
queue := queue.New()
```

Add an item to the Queue:

```go
queue.Push(42)
```

Remove an item from the Queue:

```go
queue.Pop()
```

Look at the values at the Front or Back of the Queue:

```go
queue.Front()

queue.Back()
```
