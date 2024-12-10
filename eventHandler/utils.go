package flatbuffer

import (
	"clubsocket/eventSchemaHandler"

	flatbuffers "github.com/google/flatbuffers/go"
)

func CreateEventWrapper(builder *flatbuffers.Builder, clubIdBuffer flatbuffers.UOffsetT, event flatbuffers.UOffsetT, payloadType eventSchemaHandler.Payload, payload flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	eventSchemaHandler.EventWrapperStart(builder)
	eventSchemaHandler.EventWrapperAddClubId(builder, clubIdBuffer)
	eventSchemaHandler.EventWrapperAddEvent(builder, event)
	eventSchemaHandler.EventWrapperAddPayloadType(builder, payloadType)
	eventSchemaHandler.EventWrapperAddPayload(builder, payload)

	return eventSchemaHandler.EventWrapperEnd(builder)
}
