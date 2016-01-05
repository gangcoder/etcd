// Copyright 2016 CoreOS, Inc.
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

package v3rpc

import (
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/etcdserver"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

type LeaseServer struct {
	le etcdserver.Lessor
}

func NewLeaseServer(le etcdserver.Lessor) pb.LeaseServer {
	return &LeaseServer{le: le}
}

func (ls *LeaseServer) LeaseCreate(ctx context.Context, cr *pb.LeaseCreateRequest) (*pb.LeaseCreateResponse, error) {
	return ls.le.LeaseCreate(ctx, cr)
}

func (ls *LeaseServer) LeaseRevoke(ctx context.Context, rr *pb.LeaseRevokeRequest) (*pb.LeaseRevokeResponse, error) {
	r, err := ls.le.LeaseRevoke(ctx, rr)
	if err != nil {
		return nil, ErrLeaseNotFound
	}
	return r, nil
}

func (ls *LeaseServer) LeaseKeepAlive(stream pb.Lease_LeaseKeepAliveServer) error {
	panic("not implemented")
}
