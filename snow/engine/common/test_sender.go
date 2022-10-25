// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	_ Sender = &SenderTest{}

	errAccept                = errors.New("unexpectedly called Accept")
	errSendAppRequest        = errors.New("unexpectedly called SendAppRequest")
	errSendAppResponse       = errors.New("unexpectedly called SendAppResponse")
	errSendAppGossip         = errors.New("unexpectedly called SendAppGossip")
	errSendAppGossipSpecific = errors.New("unexpectedly called SendAppGossipSpecific")
)

// SenderTest is a test sender
type SenderTest struct {
	T *testing.T

	CantAccept,
	CantSendGetStateSummaryFrontier, CantSendStateSummaryFrontier,
	CantSendGetAcceptedStateSummary, CantSendAcceptedStateSummary,
	CantSendGetAcceptedFrontier, CantSendAcceptedFrontier,
	CantSendGetAccepted, CantSendAccepted,
	CantSendGet, CantSendGetAncestors, CantSendPut, CantSendAncestors,
	CantSendPullQuery, CantSendPushQuery, CantSendChits,
	CantSendGossip,
	CantSendAppRequest, CantSendAppResponse, CantSendAppGossip, CantSendAppGossipSpecific,
	CantSendCrossChainAppRequest, CantSendCrossChainAppResponse bool

	AcceptF                      func(*snow.ConsensusContext, ids.ID, []byte) error
	SendGetStateSummaryFrontierF func(set.Set[ids.NodeID], uint32)
	SendStateSummaryFrontierF    func(ids.NodeID, uint32, []byte)
	SendGetAcceptedStateSummaryF func(set.Set[ids.NodeID], uint32, []uint64)
	SendAcceptedStateSummaryF    func(ids.NodeID, uint32, []ids.ID)
	SendGetAcceptedFrontierF     func(set.Set[ids.NodeID], uint32)
	SendAcceptedFrontierF        func(ids.NodeID, uint32, []ids.ID)
	SendGetAcceptedF             func(set.Set[ids.NodeID], uint32, []ids.ID)
	SendAcceptedF                func(ids.NodeID, uint32, []ids.ID)
	SendGetF                     func(ids.NodeID, uint32, ids.ID)
	SendGetAncestorsF            func(ids.NodeID, uint32, ids.ID)
	SendPutF                     func(ids.NodeID, uint32, []byte)
	SendAncestorsF               func(ids.NodeID, uint32, [][]byte)
	SendPushQueryF               func(set.Set[ids.NodeID], uint32, []byte)
	SendPullQueryF               func(set.Set[ids.NodeID], uint32, ids.ID)
	SendChitsF                   func(ids.NodeID, uint32, []ids.ID)
	SendGossipF                  func([]byte)
	SendAppRequestF              func(set.Set[ids.NodeID], uint32, []byte) error
	SendAppResponseF             func(ids.NodeID, uint32, []byte) error
	SendAppGossipF               func([]byte) error
	SendAppGossipSpecificF       func(set.Set[ids.NodeID], []byte) error
	SendCrossChainAppRequestF    func(ids.ID, uint32, []byte)
	SendCrossChainAppResponseF   func(ids.ID, uint32, []byte)
}

// Default set the default callable value to [cant]
func (s *SenderTest) Default(cant bool) {
	s.CantAccept = cant
	s.CantSendGetStateSummaryFrontier = cant
	s.CantSendStateSummaryFrontier = cant
	s.CantSendGetAcceptedStateSummary = cant
	s.CantSendAcceptedStateSummary = cant
	s.CantSendGetAcceptedFrontier = cant
	s.CantSendAcceptedFrontier = cant
	s.CantSendGetAccepted = cant
	s.CantSendAccepted = cant
	s.CantSendGet = cant
	s.CantSendGetAccepted = cant
	s.CantSendPut = cant
	s.CantSendAncestors = cant
	s.CantSendPullQuery = cant
	s.CantSendPushQuery = cant
	s.CantSendChits = cant
	s.CantSendGossip = cant
	s.CantSendAppRequest = cant
	s.CantSendAppResponse = cant
	s.CantSendAppGossip = cant
	s.CantSendAppGossipSpecific = cant
	s.CantSendCrossChainAppRequest = cant
	s.CantSendCrossChainAppResponse = cant
}

// SendGetStateSummaryFrontier calls SendGetStateSummaryFrontierF if it was initialized. If it
// wasn't initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) Accept(ctx *snow.ConsensusContext, containerID ids.ID, container []byte) error {
	if s.AcceptF != nil {
		return s.AcceptF(ctx, containerID, container)
	}
	if !s.CantAccept {
		return nil
	}
	if s.T != nil {
		s.T.Fatal(errAccept)
	}
	return errAccept
}

// SendGetStateSummaryFrontier calls SendGetStateSummaryFrontierF if it was initialized. If it
// wasn't initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendGetStateSummaryFrontier(validatorIDs set.Set[ids.NodeID], requestID uint32) {
	if s.SendGetStateSummaryFrontierF != nil {
		s.SendGetStateSummaryFrontierF(validatorIDs, requestID)
	} else if s.CantSendGetStateSummaryFrontier && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendGetStateSummaryFrontier")
	}
}

// SendAcceptedFrontier calls SendAcceptedFrontierF if it was initialized. If it
// wasn't initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendStateSummaryFrontier(validatorID ids.NodeID, requestID uint32, summary []byte) {
	if s.SendStateSummaryFrontierF != nil {
		s.SendStateSummaryFrontierF(validatorID, requestID, summary)
	} else if s.CantSendStateSummaryFrontier && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendStateSummaryFrontier")
	}
}

// SendGetAcceptedStateSummary calls SendGetAcceptedStateSummaryF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendGetAcceptedStateSummary(nodeIDs set.Set[ids.NodeID], requestID uint32, heights []uint64) {
	if s.SendGetAcceptedStateSummaryF != nil {
		s.SendGetAcceptedStateSummaryF(nodeIDs, requestID, heights)
	} else if s.CantSendGetAcceptedStateSummary && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendGetAcceptedStateSummaryF")
	}
}

// SendAcceptedStateSummary calls SendAcceptedStateSummaryF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendAcceptedStateSummary(validatorID ids.NodeID, requestID uint32, summaryIDs []ids.ID) {
	if s.SendAcceptedStateSummaryF != nil {
		s.SendAcceptedStateSummaryF(validatorID, requestID, summaryIDs)
	} else if s.CantSendAcceptedStateSummary && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendAcceptedStateSummary")
	}
}

// SendGetAcceptedFrontier calls SendGetAcceptedFrontierF if it was initialized.
// If it wasn't initialized and this function shouldn't be called and testing
// was initialized, then testing will fail.
func (s *SenderTest) SendGetAcceptedFrontier(validatorIDs set.Set[ids.NodeID], requestID uint32) {
	if s.SendGetAcceptedFrontierF != nil {
		s.SendGetAcceptedFrontierF(validatorIDs, requestID)
	} else if s.CantSendGetAcceptedFrontier && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendGetAcceptedFrontier")
	}
}

// SendAcceptedFrontier calls SendAcceptedFrontierF if it was initialized. If it
// wasn't initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendAcceptedFrontier(validatorID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	if s.SendAcceptedFrontierF != nil {
		s.SendAcceptedFrontierF(validatorID, requestID, containerIDs)
	} else if s.CantSendAcceptedFrontier && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendAcceptedFrontier")
	}
}

// SendGetAccepted calls SendGetAcceptedF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendGetAccepted(nodeIDs set.Set[ids.NodeID], requestID uint32, containerIDs []ids.ID) {
	if s.SendGetAcceptedF != nil {
		s.SendGetAcceptedF(nodeIDs, requestID, containerIDs)
	} else if s.CantSendGetAccepted && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendGetAccepted")
	}
}

// SendAccepted calls SendAcceptedF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendAccepted(validatorID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	if s.SendAcceptedF != nil {
		s.SendAcceptedF(validatorID, requestID, containerIDs)
	} else if s.CantSendAccepted && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendAccepted")
	}
}

// SendGet calls SendGetF if it was initialized. If it wasn't initialized and
// this function shouldn't be called and testing was initialized, then testing
// will fail.
func (s *SenderTest) SendGet(vdr ids.NodeID, requestID uint32, vtxID ids.ID) {
	if s.SendGetF != nil {
		s.SendGetF(vdr, requestID, vtxID)
	} else if s.CantSendGet && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendGet")
	}
}

// SendGetAncestors calls SendGetAncestorsF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendGetAncestors(validatorID ids.NodeID, requestID uint32, vtxID ids.ID) {
	if s.SendGetAncestorsF != nil {
		s.SendGetAncestorsF(validatorID, requestID, vtxID)
	} else if s.CantSendGetAncestors && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendCantSendGetAncestors")
	}
}

// SendPut calls SendPutF if it was initialized. If it wasn't initialized and
// this function shouldn't be called and testing was initialized, then testing
// will fail.
func (s *SenderTest) SendPut(vdr ids.NodeID, requestID uint32, vtx []byte) {
	if s.SendPutF != nil {
		s.SendPutF(vdr, requestID, vtx)
	} else if s.CantSendPut && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendPut")
	}
}

// SendAncestors calls SendAncestorsF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendAncestors(vdr ids.NodeID, requestID uint32, vtxs [][]byte) {
	if s.SendAncestorsF != nil {
		s.SendAncestorsF(vdr, requestID, vtxs)
	} else if s.CantSendAncestors && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendAncestors")
	}
}

// SendPushQuery calls SendPushQueryF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendPushQuery(vdrs set.Set[ids.NodeID], requestID uint32, vtx []byte) {
	if s.SendPushQueryF != nil {
		s.SendPushQueryF(vdrs, requestID, vtx)
	} else if s.CantSendPushQuery && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendPushQuery")
	}
}

// SendPullQuery calls SendPullQueryF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendPullQuery(vdrs set.Set[ids.NodeID], requestID uint32, vtxID ids.ID) {
	if s.SendPullQueryF != nil {
		s.SendPullQueryF(vdrs, requestID, vtxID)
	} else if s.CantSendPullQuery && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendPullQuery")
	}
}

// SendChits calls SendChitsF if it was initialized. If it wasn't initialized
// and this function shouldn't be called and testing was initialized, then
// testing will fail.
func (s *SenderTest) SendChits(vdr ids.NodeID, requestID uint32, votes []ids.ID) {
	if s.SendChitsF != nil {
		s.SendChitsF(vdr, requestID, votes)
	} else if s.CantSendChits && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendChits")
	}
}

// SendGossip calls SendGossipF if it was initialized. If it wasn't initialized
// and this function shouldn't be called and testing was initialized, then
// testing will fail.
func (s *SenderTest) SendGossip(container []byte) {
	if s.SendGossipF != nil {
		s.SendGossipF(container)
	} else if s.CantSendGossip && s.T != nil {
		s.T.Fatalf("Unexpectedly called SendGossip")
	}
}

func (s *SenderTest) SendCrossChainAppRequest(chainID ids.ID, requestID uint32, appRequestBytes []byte) error {
	if s.SendCrossChainAppRequestF != nil {
		s.SendCrossChainAppRequestF(chainID, requestID, appRequestBytes)
	} else if s.CantSendCrossChainAppRequest && s.T != nil {
		s.T.Fatal("Unexpectedly called SendCrossChainAppRequest")
	}
	return nil
}

func (s *SenderTest) SendCrossChainAppResponse(chainID ids.ID, requestID uint32, appResponseBytes []byte) error {
	if s.SendCrossChainAppResponseF != nil {
		s.SendCrossChainAppResponseF(chainID, requestID, appResponseBytes)
	} else if s.CantSendCrossChainAppResponse && s.T != nil {
		s.T.Fatal("Unexpectedly called SendCrossChainAppResponse")
	}
	return nil
}

// SendAppRequest calls SendAppRequestF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendAppRequest(nodeIDs set.Set[ids.NodeID], requestID uint32, appRequestBytes []byte) error {
	switch {
	case s.SendAppRequestF != nil:
		return s.SendAppRequestF(nodeIDs, requestID, appRequestBytes)
	case s.CantSendAppRequest && s.T != nil:
		s.T.Fatal(errSendAppRequest)
	}
	return errSendAppRequest
}

// SendAppResponse calls SendAppResponseF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendAppResponse(nodeID ids.NodeID, requestID uint32, appResponseBytes []byte) error {
	switch {
	case s.SendAppResponseF != nil:
		return s.SendAppResponseF(nodeID, requestID, appResponseBytes)
	case s.CantSendAppResponse && s.T != nil:
		s.T.Fatal(errSendAppResponse)
	}
	return errSendAppResponse
}

// SendAppGossip calls SendAppGossipF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendAppGossip(appGossipBytes []byte) error {
	switch {
	case s.SendAppGossipF != nil:
		return s.SendAppGossipF(appGossipBytes)
	case s.CantSendAppGossip && s.T != nil:
		s.T.Fatal(errSendAppGossip)
	}
	return errSendAppGossip
}

// SendAppGossipSpecific calls SendAppGossipSpecificF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *SenderTest) SendAppGossipSpecific(nodeIDs set.Set[ids.NodeID], appGossipBytes []byte) error {
	switch {
	case s.SendAppGossipSpecificF != nil:
		return s.SendAppGossipSpecificF(nodeIDs, appGossipBytes)
	case s.CantSendAppGossipSpecific && s.T != nil:
		s.T.Fatal(errSendAppGossipSpecific)
	}
	return errSendAppGossipSpecific
}
