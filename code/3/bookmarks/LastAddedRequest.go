// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bookmarks

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LastAddedRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsLastAddedRequest(buf []byte, offset flatbuffers.UOffsetT) *LastAddedRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LastAddedRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *LastAddedRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LastAddedRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func LastAddedRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func LastAddedRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
