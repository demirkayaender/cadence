// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package timeout

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"
	"time"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/client/matching"
	"github.com/uber/cadence/common/types"
)

var _ matching.Client = (*matchingClient)(nil)

// matchingClient implements the matching.Client interface instrumented with timeouts
type matchingClient struct {
	client          matching.Client
	longPollTimeout time.Duration
	timeout         time.Duration
}

// NewMatchingClient creates a new matchingClient instance
func NewMatchingClient(
	client matching.Client,
	longPollTimeout time.Duration,
	timeout time.Duration,
) matching.Client {
	return &matchingClient{
		client:          client,
		longPollTimeout: longPollTimeout,
		timeout:         timeout,
	}
}

func (c *matchingClient) AddActivityTask(ctx context.Context, ap1 *types.AddActivityTaskRequest, p1 ...yarpc.CallOption) (ap2 *types.AddActivityTaskResponse, err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.AddActivityTask(ctx, ap1, p1...)
}

func (c *matchingClient) AddDecisionTask(ctx context.Context, ap1 *types.AddDecisionTaskRequest, p1 ...yarpc.CallOption) (ap2 *types.AddDecisionTaskResponse, err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.AddDecisionTask(ctx, ap1, p1...)
}

func (c *matchingClient) CancelOutstandingPoll(ctx context.Context, cp1 *types.CancelOutstandingPollRequest, p1 ...yarpc.CallOption) (err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.CancelOutstandingPoll(ctx, cp1, p1...)
}

func (c *matchingClient) DescribeTaskList(ctx context.Context, mp1 *types.MatchingDescribeTaskListRequest, p1 ...yarpc.CallOption) (dp1 *types.DescribeTaskListResponse, err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.DescribeTaskList(ctx, mp1, p1...)
}

func (c *matchingClient) GetTaskListsByDomain(ctx context.Context, gp1 *types.GetTaskListsByDomainRequest, p1 ...yarpc.CallOption) (gp2 *types.GetTaskListsByDomainResponse, err error) {
	return c.client.GetTaskListsByDomain(ctx, gp1, p1...)
}

func (c *matchingClient) ListTaskListPartitions(ctx context.Context, mp1 *types.MatchingListTaskListPartitionsRequest, p1 ...yarpc.CallOption) (lp1 *types.ListTaskListPartitionsResponse, err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.ListTaskListPartitions(ctx, mp1, p1...)
}

func (c *matchingClient) PollForActivityTask(ctx context.Context, mp1 *types.MatchingPollForActivityTaskRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingPollForActivityTaskResponse, err error) {
	ctx, cancel := createContext(ctx, c.longPollTimeout)
	defer cancel()
	return c.client.PollForActivityTask(ctx, mp1, p1...)
}

func (c *matchingClient) PollForDecisionTask(ctx context.Context, mp1 *types.MatchingPollForDecisionTaskRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingPollForDecisionTaskResponse, err error) {
	ctx, cancel := createContext(ctx, c.longPollTimeout)
	defer cancel()
	return c.client.PollForDecisionTask(ctx, mp1, p1...)
}

func (c *matchingClient) QueryWorkflow(ctx context.Context, mp1 *types.MatchingQueryWorkflowRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingQueryWorkflowResponse, err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.QueryWorkflow(ctx, mp1, p1...)
}

func (c *matchingClient) RefreshTaskListPartitionConfig(ctx context.Context, mp1 *types.MatchingRefreshTaskListPartitionConfigRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingRefreshTaskListPartitionConfigResponse, err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.RefreshTaskListPartitionConfig(ctx, mp1, p1...)
}

func (c *matchingClient) RespondQueryTaskCompleted(ctx context.Context, mp1 *types.MatchingRespondQueryTaskCompletedRequest, p1 ...yarpc.CallOption) (err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.RespondQueryTaskCompleted(ctx, mp1, p1...)
}

func (c *matchingClient) UpdateTaskListPartitionConfig(ctx context.Context, mp1 *types.MatchingUpdateTaskListPartitionConfigRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingUpdateTaskListPartitionConfigResponse, err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.UpdateTaskListPartitionConfig(ctx, mp1, p1...)
}
