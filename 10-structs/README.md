# ✅ Quest 10 — Structs

## What are Structs?

User-defined types that group related fields. The foundation of Go's type system. Methods attach behavior to structs. **Embedding** composes types — one struct inside another inherits the inner struct's fields and methods (no classical inheritance, just composition).

```go
type Person struct {
    Name string
    Age  int
}
// func (receiver) func_name (paramenters) type {}
func (p Person) Greet() string { return "Hi, " + p.Name }

// a "method" is a func that is tied to the struct by a "receiver" which is the struct type itself or its pointer
```

## The Challenge

Model a mini **library system**:

1. Create an `Author` struct with `Name string` and `Country string`
2. Create a `Book` struct that **embeds** `Author` and has `Title string`, `Year int`, `Pages int`
3. Add a `String() string` method to `Book` that returns a formatted summary
4. Create a `Library` struct with a `[]Book` slice and two methods:
   - `AddBook(b Book)` — appends to the collection
   - `FindByAuthor(name string) []Book` — returns all books by that author
5. In main, add 4 books and search for an author who has 2 books

## What you will practice

- Struct declaration and initialization
- Embedding (promoted fields and methods)
- Methods with value and pointer receivers
- `fmt.Stringer` interface (`String() string`)
- Slice of structs

## Hints

- Embedded struct fields are promoted: `book.Name` works if `Author` is embedded
- Initialize with struct literal: `Book{Author: Author{Name: "..."}, Title: "..."}`
- `String()` lets `fmt.Println(book)` print your custom format automatically
