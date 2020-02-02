# errors
Package errors provides simple error handling primitives.

```
go get github.com/matiasvarela/errors
````

## Defining errors types
First of all, you have to define what types of errors you are going to use in your entire application. This could be done as global variables somewhere in the project.

```go
Internal := errors.Define("internal")
NotFound := errors.Define("not_found)
InvalidInput := errors.Define("invalid_input")
```

## Creating new error
The funcion errors.New create a new error based on the previously defined errors.

```go
// GetAnimal retrieves the animal with the given id, or nil if it does not exists.
func GetAnimal(id string) (Animal, error) {
    animal, err := animals.GetByID(id)
    if err != nil {
        return errors.New(Internal, err, "get animal from storage has failed")
    }

    if animal == nil {
        return errors.New(NotFound, nil, "animal has not been found")
    }

    return animal, nil
}
```

## Wrapping an error
The function errors.Wrap wraps an error with a new message but keeping the same error type.

```go
func GetAnimalName(id string) (string, error) {
    animal, err := GetAnimal(id)
    if err != nil {
        return errors.Wrap(err, "get animal has failed")
    }

    return animal.Name, nil 
}
```

## Print an error
The function errors.String returns a string containing all the information of the error including the stacktrace.

```go
func main() {
    name, err := GetAnimalName("woofy")
    if err != nil {
        println(errors.String(err))
        panic(err)
    }

    println(name)
}
```

If the animal name is not found, then the following output will be printed in the console.

```
get animal has failed | CODE: not_found | FILE: ../main.go:5 | CAUSE: {animal has not been found | CODE: not_found | FILE ../animals.go:10}
```

    In this example the files path has been simplified, but the whole path will be output.



## Contact
mtsvrl@gmail.com







