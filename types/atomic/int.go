// Copyright 2013-2014 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package atomic

import (
	"sync"
)

// AtomicInt implements an int value with atomic semantics
type AtomicInt struct {
	val   int
	mutex sync.RWMutex
}

// NewAtomicInt generates a new AtomicInt instance.
func NewAtomicInt(value int) *AtomicInt {
	return &AtomicInt{
		val: value,
	}
}

// AddAndGet atomically adds the given value to the current value.
func (ai *AtomicInt) AddAndGet(delta int) int {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	ai.val += delta
	return ai.val
}

// CompareAndSet atomically sets the value to the given updated value if the current value == expected value.
// Returns true if the expectation was met
func (ai *AtomicInt) CompareAndSet(expect int, update int) bool {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	if ai.val == expect {
		ai.val = update
		return true
	}
	return false
}

// DecrementAndGet atomically decrements current value by one and returns the result.
func (ai *AtomicInt) DecrementAndGet() int {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	ai.val--
	return ai.val
}

// Get atomically retrieves the current value.
func (ai *AtomicInt) Get() int {
	ai.mutex.RLock()
	defer ai.mutex.RUnlock()
	return ai.val
}

// GetAndAdd atomically adds the given delta to the current value and returns the result.
func (ai *AtomicInt) GetAndAdd(delta int) int {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	old := ai.val
	ai.val += delta
	return old
}

// GetAndDecrement atomically decrements the current value by one and returns the result.
func (ai *AtomicInt) GetAndDecrement() int {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	old := ai.val
	ai.val--
	return old
}

// GetAndIncrement atomically increments current value by one and returns the result.
func (ai *AtomicInt) GetAndIncrement() int {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	old := ai.val
	ai.val++
	return old
}

// GetAndSet atomically sets current value to the given value and returns the old value.
func (ai *AtomicInt) GetAndSet(newValue int) int {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	old := ai.val
	ai.val = newValue
	return old
}

// IncrementAndGet atomically increments current value by one and returns the result.
func (ai *AtomicInt) IncrementAndGet() int {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	ai.val++
	return ai.val
}

// Set atomically sets current value to the given value.
func (ai *AtomicInt) Set(newValue int) {
	ai.mutex.Lock()
	defer ai.mutex.Unlock()
	ai.val = newValue
}
