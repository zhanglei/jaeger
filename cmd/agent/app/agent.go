// Copyright (c) 2017 Uber Technologies, Inc.
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

package app

import (
	"io"
	"net"
	"net/http"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/cmd/agent/app/processors"
)

// Agent服务包括两个作用：
// 1. Sampling Server, 获取采样策略
// 2. workers goroutine，当udp server获取到本地host上的client发送的trace包后，作为生产者把trace包发送到channel队列， 然后再workers作为消费者，从channel队列上抢占trace包, 最后通过handle的Process把trace包push给Collector server
type Agent struct {
	// workers工作进程用于获取
	processors []processors.Processor
	httpServer *http.Server
	httpAddr   atomic.Value // string, set once agent starts listening
	logger     *zap.Logger
	closer     io.Closer
}

// NewAgent creates the new Agent.
func NewAgent(
	processors []processors.Processor,
	httpServer *http.Server,
	logger *zap.Logger,
) *Agent {
	a := &Agent{
		processors: processors,
		httpServer: httpServer,
		logger:     logger,
	}
	a.httpAddr.Store("")
	return a
}

// Run方法启动tcp server，同时启动worker goroutines
func (a *Agent) Run() error {
	listener, err := net.Listen("tcp", a.httpServer.Addr)
	if err != nil {
		return err
	}
	a.httpAddr.Store(listener.Addr().String())
	a.closer = listener
	go func() {
		if err := a.httpServer.Serve(listener); err != nil {
			a.logger.Error("http server failure", zap.Error(err))
		}
		a.logger.Info("agent's http server exiting")
	}()
	for _, processor := range a.processors {
		go processor.Serve()
	}
	return nil
}

// HTTPAddr returns the address that HTTP server is listening on
func (a *Agent) HTTPAddr() string {
	return a.httpAddr.Load().(string)
}

// Stop forces all agent go routines to exit.
func (a *Agent) Stop() {
	for _, processor := range a.processors {
		go processor.Stop()
	}
	a.closer.Close()
}