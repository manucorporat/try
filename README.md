#Try/catch/finally in Go
This is an experiment in Golang, that tries to bring the exception behaviour of java/python/c++ to Golang.

##Different cases of try/catch/finally
Pseudocode


###1. No crash at all

```java
try {
	print "1"
} catch(err) {
	print "2"
	print err
} finally {
	print "3"
}
print "4"
```
prints
```
1
3
4
```

###2. Throw in try

```java
try {
	print "1"
	throw "exception 1"
} catch(err) {
	print "2"
	print err
} finally {
	print "3"
}
print "4"
```
prints
```
1
2
exception 1
3
4
```

###3. Throw in try & catch

```java
try {
	print "1"
	throw "exception 1"
} catch(err) {
	print "2"
	throw "exception 2"
} finally {
	print "3"
}
print "4"
```
prints
```
1
2
3
---> uncatched exception 2
```

###4. Throw in try & catch & finally

```java
try {
	print "1"
	throw "exception 1"
} catch(err) {
	print "2"
	throw "exception 2"
} finally {
	print "3"
	throw "exception 3"
}
print "4"
```
prints
```
1
2
3
---> uncatched exception 3
```
yes! "exception 2" was throwed but "overwritten" by "exception 3"

###5. Finally is always optional

```java
try {
	print "1"
	throw "exception 1"
} catch(err) {
	print "2"
	throw "exception 2"
}
print "4"
```
prints
```
1
2
---> uncatched exception 2
```

###5. Rethrowing exceptions

```java
try {
	print "1"
	throw "exception 1"
} catch() {
	print "2"
	throw
}
print "4"
```
prints
```
1
2
---> uncatched exception 1
```

##Porting all this to Go

1. We need a Try, Catch and Finally method
2. We need a Throw() method for rethrowing
3. It needs to be state-less

###Examples:
####1. Simple panic inside try
```go
import (
	"fmt"
	"github.com/manucorporat/try"
)

func main() {
	var obj interface{}
	obj = 2
	try.This(func() {
		// this operation will panic because obj is an integer
		text := obj.(string)
		fmt.Println(text)

	}).Catch(func(e try.E) {
		// Print crash
		fmt.Println(e)
	})
}
```

####2. Add Finally
Unfortunaly we have to include Finally before Catch.  Though the behaviour is exactly the same than java.

```go
import (
	"fmt"
	"github.com/manucorporat/try"
)

func main() {
	try.This(func() {
		panic("my panic")

	}).Finally(func() {
		fmt.Println("this must be printed after the catch")

	}).Catch(func(e try.E) {
		// Print crash
		fmt.Println(e)
	})
}
```

####3. Rethrowing
Unfortunaly we have to include Finally before Catch.  Though the behaviour is exactly the same than java.

```go
import (
	"fmt"
	"github.com/manucorporat/try"
)

func main() {
	try.This(func() {
		panic("my panic")

	}).Finally(func() {
		fmt.Println("this must be printed after the catch")

	}).Catch(func(e try.E) {
		fmt.Println("exception catched") // print
		try.Throw()                      // rethrow current exception!!
	})
}
```

####4. Nested

```go
package main

import (
	"fmt"
	"github.com/manucorporat/try"
)

func main() {
	try.This(func() {
		try.This(func() {
			panic("my panic")

		}).Catch(func(e try.E) {
			fmt.Println("fixing stuff") // print
			try.Throw()                 // rethrow current exception!!
		})

	}).Catch(func(e try.E) {
		// print
		fmt.Println(e)
	})
	fmt.Println("hey")
}
```
prints
```
fixing stuff
my panic
hey
```

### Full covered with unit tests
See test_try.go


