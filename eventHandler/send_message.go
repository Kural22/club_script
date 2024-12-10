package flatbuffer

import (
	"clubsocket/eventSchemaHandler"

	flatbuffers "github.com/google/flatbuffers/go"
)

func SendMessageEvent(clubId string, msgToSend string, msgType string) []byte {
	// Get the builder and event string
	eventToSend := "SEND_MESSAGE"
	// builder, event := getBuilderWithEvent(clubId, eventToSend, msgToSend, msgType)
	builder := flatbuffers.NewBuilder(1024)
	clubIdBuffer := builder.CreateString(clubId)

	// Create the message and type strings in the same builder
	event := builder.CreateString(eventToSend)
	message := builder.CreateString(msgToSend)
	typeStr := builder.CreateString(msgType)

	// Create the SendMessagePayload in the builder
	messagePayload := CreateSendMessagePayload(builder, message, typeStr)

	msg := CreateEventWrapper(builder, clubIdBuffer, event, eventSchemaHandler.PayloadSendMessagePayload, messagePayload)

	builder.Finish(msg)

	return builder.FinishedBytes()
}

func CreateSendMessagePayload(builder *flatbuffers.Builder, msgOffset flatbuffers.UOffsetT, msgTypeOffset flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	// Start building the SendMessagePayload
	eventSchemaHandler.SendMessagePayloadStart(builder)

	// Add the message and message type fields
	eventSchemaHandler.SendMessagePayloadAddMsg(builder, msgOffset)
	eventSchemaHandler.SendMessagePayloadAddMsgType(builder, msgTypeOffset)

	// End the creation of the SendMessagePayload table and return the offset
	return eventSchemaHandler.SendMessagePayloadEnd(builder)
}
