# fu: functional utilities

If you miss your maps, reducers and filters from other languages, and don't mind manually unwrapping `interface{}`s, this library is for you.

## Mapping

```go
ctx := context.Background()
ns := []interface{}{1, 2, 3, 4, 5}
ms, err := fu.MapFn(ctx, ns, func(ctx context.Context, i interface{}) (interface{}, error){
  n, ok := i.(int)
  if !ok {
    return errors.New("invalid type")
  }
  return n + 1
})
if err != nil {
  return err
}
fmt.Printf("%v\n", ms) // []interface{}{2, 3, 4, 5, 6}
```

Or as a shorthand:

```go
ctx := context.Background()
ns := []interface{}{1, 2, 3, 4, 5}
ms, err := fu.Map(ctx, ns, fu.Add(1))
if err != nil {
  return err
}
fmt.Printf("%v\n", ms) // []interface{}{2, 3, 4, 5, 6}
```

And if you want to do stuff in parallel:

```go
ctx := context.Background()
ns := []interface{}{1, 2, 3, 4, 5}
ms, err := fu.ParallelMap(ctx, ns, fu.Add(1))
if err != nil {
  return err
}
fmt.Printf("%v\n", ms) // []interface{}{2, 3, 4, 5, 6}
```