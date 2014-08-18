[![GoDoc](https://godoc.org/github.com/manucorporat/try?status.png)](https://godoc.org/github.com/manucorporat/try)  
#Try/Catch/Finally in Go
Experiment in Golang that tries to bring the exception behaviour of Java/Python/C++ to Golang using the same syntax.  

## Experiment only!!
I as a Go developer, do not recomend to use this library in production or real world code.
The Go language was designed to do not use expections, instead use the explicit error management suggested in [Effective Go](http://golang.org/doc/effective_go.html). A good programmer MUST write idiomatic code.

Instead you should use this project as a learning tool to understand the exceptions flow in a language like python/c++ and python.

It also shows that Go, even without explicit exceptions has the semantics needed to provide it exactly in the same way Java/C++ and Python does.


##Approach

1. We need `Try`, `Catch` and `Finally` methods.
2. We need a `Throw()` method for rethrowing exceptions.
3. It needs to be stateless so it could be nested and used across many threads.  

###API examples:  

####1. Simple `panic()` inside `Try`  

Unfortunately we have to include a `Finally` before a `Catch`. I have tried to find a way to avoid it, but looks impossible. Anyway, the behaviour and order of call is exactly the same than Java or Python.  

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

####2. `Finally` is optional  

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

####3. Rethrowing  

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

	}).Catch(func(_ try.E) {
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
* See [test_try.go](https://github.com/manucorporat/try/blob/master/try.go)  


##Different cases of try/catch/finally  

This Go package has the same behaviour than the implementation of exceptions in Java, C++ and Python.  

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

###2. Throw in `try`  

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

###3. Throw in `try` and `catch`  

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

###4. Throw in `try`, `catch` and `finally`  

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

###5. `finally` is optional  

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
