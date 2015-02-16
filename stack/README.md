# Stack

A simple LIFO stack implementation backed by a singly linked list.

## Import

```go
import (
    "github.com/connorjacobsen/gorithms/containers/stack"
)
```

## API

```go
type Element struct {
    // The value of the element.
    Value interface{}
}

type Stack struct{}
```

Create a new Stack

```go
stack := stack.New()
```

Get the size of the Stack

```go
stack.Size()
```

Push a value to the Stack

```go
stack.Push(42)
```

Pop a value off the Stack

```go
data := stack.Pop()
```

Peek at the top of the Stack

```go
top := stack.Peek()
```

## TODO

I would like to make this implementation thread safe. I don't know that anybody would want to use it in that way, but I think it could be useful to some people (or myself), so I want to do it.
