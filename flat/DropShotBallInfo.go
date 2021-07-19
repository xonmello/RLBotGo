// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DropShotBallInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsDropShotBallInfo(buf []byte, offset flatbuffers.UOffsetT) *DropShotBallInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DropShotBallInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DropShotBallInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DropShotBallInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DropShotBallInfo) AbsorbedForce() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *DropShotBallInfo) MutateAbsorbedForce(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func (rcv *DropShotBallInfo) DamageIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DropShotBallInfo) MutateDamageIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *DropShotBallInfo) ForceAccumRecent() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *DropShotBallInfo) MutateForceAccumRecent(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func DropShotBallInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func DropShotBallInfoAddAbsorbedForce(builder *flatbuffers.Builder, absorbedForce float32) {
	builder.PrependFloat32Slot(0, absorbedForce, 0.0)
}
func DropShotBallInfoAddDamageIndex(builder *flatbuffers.Builder, damageIndex int32) {
	builder.PrependInt32Slot(1, damageIndex, 0)
}
func DropShotBallInfoAddForceAccumRecent(builder *flatbuffers.Builder, forceAccumRecent float32) {
	builder.PrependFloat32Slot(2, forceAccumRecent, 0.0)
}
func DropShotBallInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}