// Copyright 2016 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scmp

import (
	"fmt"

	"gopkg.in/restruct.v1"

	"github.com/scionproto/scion/go/lib/common"
	"github.com/scionproto/scion/go/lib/ctrl/path_mgmt"
	"github.com/scionproto/scion/go/lib/serrors"
	"github.com/scionproto/scion/go/lib/util"
)

type Info interface {
	fmt.Stringer
	Copy() Info
	Len() int
	Write(b common.RawBytes) (int, error)
}

var _ Info = (*InfoString)(nil)

type InfoString string

func (s InfoString) Copy() Info {
	// Strings are immutable, so no need to actually copy.
	return s
}

func (s InfoString) Len() int {
	l := 2 + len(s)
	return l + util.CalcPadding(l, common.LineLen)
}

func (s InfoString) Write(b common.RawBytes) (int, error) {
	common.Order.PutUint16(b, uint16(len(s)))
	copy(b[:2], s)
	return util.FillPadding(b, 2+len(s), common.LineLen), nil
}

func (s InfoString) String() string {
	return string(s)
}

var _ Info = (*InfoEcho)(nil)

type InfoEcho struct {
	Id  uint64
	Seq uint16
}

func InfoEchoFromRaw(b common.RawBytes) (*InfoEcho, error) {
	e := &InfoEcho{}
	if len(b) < e.Len() {
		return nil, common.NewBasicError("Can't parse SCMP Info Echo, buffer is too short",
			nil, "expected", e.Len(), "actual", len(b))
	}
	if err := restruct.Unpack(b, common.Order, e); err != nil {
		return nil, common.NewBasicError("Failed to unpack SCMP ECHO info", err)
	}
	return e, nil
}

func (e *InfoEcho) Copy() Info {
	if e == nil {
		return nil
	}
	return &InfoEcho{Id: e.Id, Seq: e.Seq}
}

func (e *InfoEcho) Len() int {
	l := 10
	return l + util.CalcPadding(l, common.LineLen)
}

func (e *InfoEcho) Write(b common.RawBytes) (int, error) {
	common.Order.PutUint64(b[0:], e.Id)
	common.Order.PutUint16(b[8:], e.Seq)
	return util.FillPadding(b, 10, common.LineLen), nil
}

func (e *InfoEcho) String() string {
	return fmt.Sprintf("Id=%v Seq=%v", e.Id, e.Seq)
}

var _ Info = (*InfoPktSize)(nil)

type InfoPktSize struct {
	Size uint16
	MTU  uint16
}

func InfoPktSizeFromRaw(b common.RawBytes) (*InfoPktSize, error) {
	p := &InfoPktSize{}
	if err := restruct.Unpack(b, common.Order, p); err != nil {
		return nil, common.NewBasicError("Failed to unpack SCMP Pkt Size info", err)
	}
	return p, nil
}

func (p *InfoPktSize) Copy() Info {
	if p == nil {
		return nil
	}
	return &InfoPktSize{Size: p.Size, MTU: p.MTU}
}

func (p *InfoPktSize) Len() int {
	l := 4
	return l + util.CalcPadding(l, common.LineLen)
}

func (p *InfoPktSize) Write(b common.RawBytes) (int, error) {
	common.Order.PutUint16(b[0:], p.Size)
	common.Order.PutUint16(b[2:], p.MTU)
	return util.FillPadding(b, 4, common.LineLen), nil
}

func (p *InfoPktSize) String() string {
	return fmt.Sprintf("Size=%v MTU=%v", p.Size, p.MTU)
}

var _ Info = (*InfoPathOffsets)(nil)

const infoPathOffsetsLen = 11

type InfoPathOffsets struct {
	InfoF   uint8
	HopF    uint8
	IfID    common.IFIDType
	Ingress bool
}

func InfoPathOffsetsFromRaw(b common.RawBytes) (*InfoPathOffsets, error) {
	p := &InfoPathOffsets{}
	if len(b) < p.Len() {
		return nil, common.NewBasicError("Can't parse InfoPathOffsets, buffer is too small", nil,
			"min", p.Len(), "actual", len(b))
	}
	// Check 0 padding
	for _, pad := range b[infoPathOffsetsLen:p.Len()] {
		if pad != 0 {
			return nil, common.NewBasicError("InfoPathOffsets padding is not zeroed", nil,
				"actual", b[infoPathOffsetsLen:p.Len()])
		}
	}
	p.InfoF = b[0]
	p.HopF = b[1]
	p.IfID = common.IFIDType(common.Order.Uint64(b[2:]))
	p.Ingress = (b[10] & 0x01) == 0x01
	return p, nil
}

func (p *InfoPathOffsets) Copy() Info {
	if p == nil {
		return nil
	}
	return &InfoPathOffsets{InfoF: p.InfoF, HopF: p.HopF, IfID: p.IfID, Ingress: p.Ingress}
}

func (p *InfoPathOffsets) Len() int {
	return infoPathOffsetsLen + util.CalcPadding(infoPathOffsetsLen, common.LineLen)
}

func (p *InfoPathOffsets) Write(b common.RawBytes) (int, error) {
	if len(b) < p.Len() {
		return 0, common.NewBasicError("Can't write InfoPathOffsets, buffer is too small", nil,
			"min", p.Len(), "actual", len(b))
	}
	b[0] = p.InfoF
	b[1] = p.HopF
	common.Order.PutUint64(b[2:], uint64(p.IfID))
	if p.Ingress {
		b[10] = 1
	} else {
		b[10] = 0
	}
	return util.FillPadding(b, 11, common.LineLen), nil
}

func (p *InfoPathOffsets) String() string {
	return fmt.Sprintf("InfoF=%d HopF=%d IfID=%d Ingress=%v", p.InfoF, p.HopF, p.IfID, p.Ingress)
}

var _ Info = (*InfoRevocation)(nil)

type InfoRevocation struct {
	*InfoPathOffsets
	RawSRev common.RawBytes
}

func NewInfoRevocation(infoF, hopF uint8, ifID common.IFIDType, ingress bool,
	rawSRev common.RawBytes) *InfoRevocation {
	return &InfoRevocation{
		InfoPathOffsets: &InfoPathOffsets{InfoF: infoF, HopF: hopF, IfID: ifID, Ingress: ingress},
		RawSRev:         rawSRev,
	}
}

func InfoRevocationFromRaw(b common.RawBytes) (*InfoRevocation, error) {
	var err error
	p := &InfoRevocation{}
	p.InfoPathOffsets, err = InfoPathOffsetsFromRaw(b)
	if err != nil {
		return nil, common.NewBasicError("Unable to parse path offsets", err)
	}
	p.RawSRev = b[p.InfoPathOffsets.Len():]
	return p, nil
}
func (r *InfoRevocation) Copy() Info {
	if r == nil {
		return nil
	}
	return &InfoRevocation{
		r.InfoPathOffsets.Copy().(*InfoPathOffsets),
		append(common.RawBytes(nil), r.RawSRev...),
	}
}

func (r *InfoRevocation) Len() int {
	l := r.InfoPathOffsets.Len() + len(r.RawSRev)
	return l + util.CalcPadding(l, common.LineLen)
}

func (r *InfoRevocation) Write(b common.RawBytes) (int, error) {
	count, err := r.InfoPathOffsets.Write(b)
	if err != nil {
		return 0, err
	}
	count += copy(b[count:], r.RawSRev)
	return util.FillPadding(b, count, common.LineLen), nil
}

func (r *InfoRevocation) String() string {
	revStr := fmt.Sprintf("InfoF=%d HopF=%d IfID=%d Ingress=%v",
		r.InfoF, r.HopF, r.IfID, r.Ingress)
	sRevInfo, err := path_mgmt.NewSignedRevInfoFromRaw(r.RawSRev)
	if err != nil {
		return fmt.Sprintf("%s RawSRev=%v", revStr, r.RawSRev)
	}
	return fmt.Sprintf("%s SRev=%s", revStr, sRevInfo)
}

var _ Info = (*InfoExtIdx)(nil)

type InfoExtIdx struct {
	Idx uint8
}

func InfoExtIdxFromRaw(b common.RawBytes) (*InfoExtIdx, error) {
	if len(b) == 0 {
		return nil, serrors.New("Unable to parse InfoExtIdx")
	}
	return &InfoExtIdx{Idx: b[0]}, nil
}

func (e *InfoExtIdx) Copy() Info {
	if e == nil {
		return nil
	}
	return &InfoExtIdx{Idx: e.Idx}
}

func (r *InfoExtIdx) Len() int {
	return 1 + util.CalcPadding(1, common.LineLen)
}

func (e *InfoExtIdx) Write(b common.RawBytes) (int, error) {
	b[0] = e.Idx
	return util.FillPadding(b, 1, common.LineLen), nil
}

func (e *InfoExtIdx) String() string {
	return fmt.Sprintf("Idx=%v", e.Idx)
}
