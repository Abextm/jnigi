# jnigi
Java Native Interface Golang Interface.

A package to access Java from Golang code. Can be used from a Golang executable or shared library.
This allows for Golang to initiate the JVM or Java to start a Golang runtime respectively.

## Notes
### Signals
The JVM calls sigaction, Golang requires that SA_ONSTACK flag be passed.
Without this Golang will not be able to print the exception information eg. type, stack trace, line number.
On Linux a solution is using LD_PRELOAD with a library that intercepts sigaction and adds the flag. (Code that does this: https://gist.github.com/timob/5d3032b54ed6ba2dc6de34b245c556c7)

## Status
* Documentation needed.
* Tests for main functions tests are present.
* Has been used in Golang (using Go 1.6) executable multi threaded applications on Linux / Windows.

## Changes
* 2016-08-01 Initial version.

## Example

```` go
package main

import (
	"fmt"
	"github.com/timob/jnigi"
	"log"
)

func main() {
	_, env, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{"-Xcheck:jni"}))
	if err != nil {
		log.Fatal(err)
	}
	obj, err := env.NewObject("java/lang/Object")
	if err != nil {
		log.Fatal(err)
	}
	v, err := obj.CallMethod(env, "hashCode", jnigi.Int)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("object hash code: %d\n", v.(int))
}

````
