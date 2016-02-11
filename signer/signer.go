// Copyright 2015 Google Inc. All Rights Reserved.
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

package signer

import (
	"log"
	"time"

	"github.com/google/e2e-key-server/builder"
	"github.com/google/e2e-key-server/db"
	"github.com/google/e2e-key-server/db/leveldb"

	proto "github.com/golang/protobuf/proto"
	corepb "github.com/google/e2e-key-server/proto/google_security_e2ekeys_core"
	v2pb "github.com/google/e2e-key-server/proto/google_security_e2ekeys_v2"
)

// Signer is the object responsible for triggering epoch creation and signing
// the epoch head once created.
type Signer struct {
	// distributed is an instance to Distributed.
	distributed db.Distributed
	// builder is signer's instance of builder.
	builder *builder.Builder
	// ticker ticks everytime a new epoch should be created.
	ticker *time.Ticker
	// local is a local store instance of the signer.
	local db.Local
}

// New creates a new instance of the signer.
func New(distributed db.Distributed, dbPath string, seconds uint) (*Signer, error) {
	local, err := leveldb.Open(dbPath)
	if err != nil {
		return nil, err
	}
	// Create the tree builder.
	b := builder.NewForSigner(distributed, local)

	// Create a signer instance.
	signer := &Signer{
		distributed: distributed,
		builder:     b,
		ticker:      time.NewTicker(time.Second * time.Duration(seconds)),
		local:       local,
	}
	go signer.createEpoch()
	return signer, nil
}

// createEpoch watches the ticker channel and triggers epoch creation once the
// ticker ticks.
func (s *Signer) createEpoch() {
	for _ = range s.ticker.C {
		lastCommitmentTS := s.builder.LastCommitmentTimestamp()
		epochHead, err := s.builder.CreateEpoch(lastCommitmentTS, true)
		if err != nil {
			log.Fatalf("Failed to create epoch: %v", err)
		}

		// Create SignedEpochHead.
		// TODO(cesarghali): fill IssueTime and PreviousEpochHeadHash.
		epochHeadData, err := proto.Marshal(epochHead)
		if err != nil {
			log.Fatalf("Failed to marshal epoch: %v", err)
		}
		signedEpochHead := &v2pb.SignedEpochHead{
			EpochHead: epochHeadData,
			// TODO(cesarghali): fill Signatures
		}

		// Write signed epoch head in the leveldb.
		epochInfo := &corepb.EpochInfo{
			SignedEpochHead:         signedEpochHead,
			LastCommitmentTimestamp: lastCommitmentTS,
		}
		if err := s.distributed.WriteEpochInfo(nil, epochHead.Epoch, epochInfo); err != nil {
			log.Fatalf("Failed to write EpochInfo: %v", err)
		}
	}
}

// Stop stops the signer and release all associated resource.
func (s *Signer) Stop() {
	s.ticker.Stop()
	s.local.Close()
	s.builder.Close()
}
