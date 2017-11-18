// Copyright 2016 Tim O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package jnigi

/*
#cgo CFLAGS:-I../include/ "-Ic:/Program Files/Java/jdk-9.0.1/include" "-Ic:/Program Files/Java/jdk-9.0.1/include/win32"
#cgo LDFLAGS:-ljvm "-Lc:/Program Files/Java/jdk-9.0.1/bin/server"
*/
import "C"
