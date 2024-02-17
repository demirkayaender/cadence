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

package thrift

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/thrift.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/common/types"
	"github.com/uber/cadence/common/types/mapper/thrift"
)

func (g matchingClient) AddActivityTask(ctx context.Context, ap1 *types.AddActivityTaskRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.AddActivityTask(ctx, thrift.FromMatchingAddActivityTaskRequest(ap1), p1...)
	return thrift.ToError(err)
}

func (g matchingClient) AddDecisionTask(ctx context.Context, ap1 *types.AddDecisionTaskRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.AddDecisionTask(ctx, thrift.FromMatchingAddDecisionTaskRequest(ap1), p1...)
	return thrift.ToError(err)
}

func (g matchingClient) CancelOutstandingPoll(ctx context.Context, cp1 *types.CancelOutstandingPollRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.CancelOutstandingPoll(ctx, thrift.FromMatchingCancelOutstandingPollRequest(cp1), p1...)
	return thrift.ToError(err)
}

func (g matchingClient) DescribeTaskList(ctx context.Context, mp1 *types.MatchingDescribeTaskListRequest, p1 ...yarpc.CallOption) (dp1 *types.DescribeTaskListResponse, err error) {
	response, err := g.c.DescribeTaskList(ctx, thrift.FromMatchingDescribeTaskListRequest(mp1), p1...)
	return thrift.ToMatchingDescribeTaskListResponse(response), thrift.ToError(err)
}

func (g matchingClient) GetTaskListsByDomain(ctx context.Context, gp1 *types.GetTaskListsByDomainRequest, p1 ...yarpc.CallOption) (gp2 *types.GetTaskListsByDomainResponse, err error) {
	response, err := g.c.GetTaskListsByDomain(ctx, thrift.FromMatchingGetTaskListsByDomainRequest(gp1), p1...)
	return thrift.ToMatchingGetTaskListsByDomainResponse(response), thrift.ToError(err)
}

func (g matchingClient) ListTaskListPartitions(ctx context.Context, mp1 *types.MatchingListTaskListPartitionsRequest, p1 ...yarpc.CallOption) (lp1 *types.ListTaskListPartitionsResponse, err error) {
	response, err := g.c.ListTaskListPartitions(ctx, thrift.FromMatchingListTaskListPartitionsRequest(mp1), p1...)
	return thrift.ToMatchingListTaskListPartitionsResponse(response), thrift.ToError(err)
}

func (g matchingClient) PollForActivityTask(ctx context.Context, mp1 *types.MatchingPollForActivityTaskRequest, p1 ...yarpc.CallOption) (pp1 *types.PollForActivityTaskResponse, err error) {
	response, err := g.c.PollForActivityTask(ctx, thrift.FromMatchingPollForActivityTaskRequest(mp1), p1...)
	return thrift.ToMatchingPollForActivityTaskResponse(response), thrift.ToError(err)
}

func (g matchingClient) PollForDecisionTask(ctx context.Context, mp1 *types.MatchingPollForDecisionTaskRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingPollForDecisionTaskResponse, err error) {
	response, err := g.c.PollForDecisionTask(ctx, thrift.FromMatchingPollForDecisionTaskRequest(mp1), p1...)
	return thrift.ToMatchingPollForDecisionTaskResponse(response), thrift.ToError(err)
}

func (g matchingClient) QueryWorkflow(ctx context.Context, mp1 *types.MatchingQueryWorkflowRequest, p1 ...yarpc.CallOption) (qp1 *types.QueryWorkflowResponse, err error) {
	response, err := g.c.QueryWorkflow(ctx, thrift.FromMatchingQueryWorkflowRequest(mp1), p1...)
	return thrift.ToMatchingQueryWorkflowResponse(response), thrift.ToError(err)
}

func (g matchingClient) RespondQueryTaskCompleted(ctx context.Context, mp1 *types.MatchingRespondQueryTaskCompletedRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondQueryTaskCompleted(ctx, thrift.FromMatchingRespondQueryTaskCompletedRequest(mp1), p1...)
	return thrift.ToError(err)
}