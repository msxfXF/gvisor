// Copyright 2019 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package test1 is a test package.
package test1

import (
	"fmt"
)

// Interface is a generic interface.
type Interface interface {
	Foo()
}

// Type is a concrete implementation of Interface.
type Type struct {
	A uint64
	B uint64
}

// Foo implements Interface.Foo.
//
//go:nosplit
func (t Type) Foo() {
	fmt.Printf("%v", t) // Never executed.
}

// InterfaceFunction is passed an interface argument.
// +checkescape:all,hard
//
//go:nosplit
func InterfaceFunction(i Interface) {
	// Do nothing; exported for tests.
}

// TypeFunction is passed a concrete pointer argument.
// +checkesacape:all,hard
//
//go:nosplit
func TypeFunction(t *Type) {
}

// BuiltinMap creates a new map.
// +mustescape:local,builtin
//
//go:nosplit
//go:noinline
func BuiltinMap(x int) map[string]bool {
	return make(map[string]bool)
}

// +mustescape:builtin
//
//go:nosplit
//go:noinline
func builtinMapRec(x int) map[string]bool {
	return BuiltinMap(x)
}

// BuiltinClosure returns a closure around x.
// +mustescape:local,builtin
//
//go:nosplit
//go:noinline
func BuiltinClosure(x int) func() {
	return func() {
		fmt.Printf("%v", x)
	}
}

// +mustescape:builtin
//
//go:nosplit
//go:noinline
func builtinClosureRec(x int) func() {
	return BuiltinClosure(x)
}

// BuiltinMakeSlice makes a new slice.
// +mustescape:local,builtin
//
//go:nosplit
//go:noinline
func BuiltinMakeSlice(x int) []byte {
	return make([]byte, x)
}

// +mustescape:builtin
//
//go:nosplit
//go:noinline
func builtinMakeSliceRec(x int) []byte {
	return BuiltinMakeSlice(x)
}

// BuiltinAppend calls append on a slice.
// +mustescape:local,builtin
//
//go:nosplit
//go:noinline
func BuiltinAppend(x []byte) []byte {
	return append(x, 0)
}

// +mustescape:builtin
//
//go:nosplit
//go:noinline
func builtinAppendRec() []byte {
	return BuiltinAppend(nil)
}

// BuiltinChan makes a channel.
// +mustescape:local,builtin
//
//go:nosplit
//go:noinline
func BuiltinChan() chan int {
	return make(chan int)
}

// +mustescape:builtin
//
//go:nosplit
//go:noinline
func builtinChanRec() chan int {
	return BuiltinChan()
}

// Heap performs an explicit heap allocation.
// +mustescape:local,heap
//
//go:nosplit
//go:noinline
func Heap() *Type {
	var t Type
	return &t
}

// +mustescape:heap
//
//go:nosplit
//go:noinline
func heapRec() *Type {
	return Heap()
}

// Dispatch dispatches via an interface.
// +mustescape:local,interface
//
//go:nosplit
//go:noinline
func Dispatch(i Interface) {
	i.Foo()
}

// +mustescape:interface
//
//go:nosplit
//go:noinline
func dispatchRec(i Interface) {
	Dispatch(i)
}

// Dynamic invokes a dynamic function.
// +mustescape:local,dynamic
//
//go:nosplit
//go:noinline
func Dynamic(f func()) {
	f()
}

// +mustescape:dynamic
//
//go:nosplit
//go:noinline
func dynamicRec(f func()) {
	Dynamic(f)
}

//
//go:nosplit
//go:noinline
func internalFunc() {
}

// Split includes a guaranteed stack split.
// +mustescape:local,stack
//
//go:noinline
func Split() {
	internalFunc()
}

// +mustescape:stack
//
//go:nosplit
//go:noinline
func splitRec() {
	Split()
}
