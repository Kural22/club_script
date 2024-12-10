package flatbuffer

import (
	"clubsocket/eventSchemaHandler"

	flatbuffers "github.com/google/flatbuffers/go"
)

func JoinClubEvent(clubId string) []byte {

	eventToSend := "SILENT_JOIN"

	builder := flatbuffers.NewBuilder(1024)
	clubIdBuffer := builder.CreateString(clubId)

	event := builder.CreateString(eventToSend)

	messagePayload := CreateJoinClubPayload(builder, clubIdBuffer)

	msg := CreateEventWrapper(builder, clubIdBuffer, event, eventSchemaHandler.PayloadJoinClubPayload, messagePayload)

	builder.Finish(msg)

	return builder.FinishedBytes()
}

func CreateJoinClubPayload(builder *flatbuffers.Builder, clubId flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	eventSchemaHandler.JoinClubPayloadStart(builder)

	eventSchemaHandler.JoinClubPayloadAddClubId(builder, clubId)

	return eventSchemaHandler.JoinClubPayloadEnd(builder)
}