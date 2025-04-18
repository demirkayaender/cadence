// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dynamicproperties

import (
	"time"
)

// These mock functions are for tests to use config properties that are dynamic

// GetIntPropertyFn returns value as IntPropertyFn
func GetIntPropertyFn(value int) func(opts ...FilterOption) int {
	return func(...FilterOption) int { return value }
}

// GetIntPropertyFilteredByDomain returns values as IntPropertyFnWithDomainFilters
func GetIntPropertyFilteredByDomain(value int) func(domain string) int {
	return func(domain string) int { return value }
}

// GetIntPropertyFilteredByTaskListInfo returns value as IntPropertyFnWithTaskListInfoFilters
func GetIntPropertyFilteredByTaskListInfo(value int) func(domain string, taskList string, taskType int) int {
	return func(domain string, taskList string, taskType int) int { return value }
}

// GetIntPropertyFilteredByShardID returns values as IntPropertyFnWithShardIDFilter
func GetIntPropertyFilteredByShardID(value int) func(shardID int) int {
	return func(shardID int) int { return value }
}

// GetIntPropertyFilteredByWorkflowType returns values as IntPropertyFnWithWorkflowTypeFilters
func GetIntPropertyFilteredByWorkflowType(value int) func(domainName string, workflowType string) int {
	return func(domainName string, workflowType string) int { return value }
}

// GetDurationPropertyFilteredByWorkflowType returns values as IntPropertyFnWithWorkflowTypeFilters
func GetDurationPropertyFilteredByWorkflowType(value time.Duration) func(domainName string, workflowType string) time.Duration {
	return func(domainName string, workflowType string) time.Duration { return value }
}

// GetFloatPropertyFn returns value as FloatPropertyFn
func GetFloatPropertyFn(value float64) func(opts ...FilterOption) float64 {
	return func(...FilterOption) float64 { return value }
}

// GetBoolPropertyFn returns value as BoolPropertyFn
func GetBoolPropertyFn(value bool) func(opts ...FilterOption) bool {
	return func(...FilterOption) bool { return value }
}

// GetBoolPropertyFnFilteredByDomain returns value as BoolPropertyFnWithDomainFilters
func GetBoolPropertyFnFilteredByDomain(value bool) func(domain string) bool {
	return func(domain string) bool { return value }
}

// GetBoolPropertyFnFilteredByDomainID returns value as BoolPropertyFnWithDomainIDFilters
func GetBoolPropertyFnFilteredByDomainID(value bool) func(domainID string) bool {
	return func(domainID string) bool { return value }
}

// GetBoolPropertyFilteredByTaskListInfo returns value as BoolPropertyFnWithTaskListInfoFilters
func GetBoolPropertyFilteredByTaskListInfo(value bool) func(domain string, taskList string, taskType int) bool {
	return func(domain string, taskList string, taskType int) bool { return value }
}

// GetDurationPropertyFnFilteredByDomain returns value as DurationPropertyFnFilteredByDomain
func GetDurationPropertyFnFilteredByDomain(value time.Duration) func(domain string) time.Duration {
	return func(domain string) time.Duration { return value }
}

// GetDurationPropertyFn returns value as DurationPropertyFn
func GetDurationPropertyFn(value time.Duration) func(opts ...FilterOption) time.Duration {
	return func(...FilterOption) time.Duration { return value }
}

// GetDurationPropertyFnFilteredByTaskListInfo returns value as DurationPropertyFnWithTaskListInfoFilters
func GetDurationPropertyFnFilteredByTaskListInfo(value time.Duration) func(domain string, taskList string, taskType int) time.Duration {
	return func(domain string, taskList string, taskType int) time.Duration { return value }
}

// GetDurationPropertyFnFilteredByShardID returns value as DurationPropertyFnWithShardIDFilter
func GetDurationPropertyFnFilteredByShardID(value time.Duration) func(shardID int) time.Duration {
	return func(shardID int) time.Duration { return value }
}

// GetStringPropertyFn returns value as StringPropertyFn
func GetStringPropertyFn(value string) func(opts ...FilterOption) string {
	return func(...FilterOption) string { return value }
}

// GetStringPropertyFnFilteredByDomain returns value as StringPropertyFnWithDomainFilters
func GetStringPropertyFnFilteredByDomain(value string) func(domain string) string {
	return func(domain string) string { return value }
}

// GetMapPropertyFn returns value as MapPropertyFn
func GetMapPropertyFn(value map[string]interface{}) func(opts ...FilterOption) map[string]interface{} {
	return func(...FilterOption) map[string]interface{} { return value }
}
