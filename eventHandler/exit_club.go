package flatbuffer

import (
	"clubsocket/eventSchemaHandler"
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

func ExitClubEvent(clubId string, userId string) []byte {

	eventToSend := "EXIT_CLUB"

	var userIds []int32

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil
	}

	userIds = append(userIds, int32(userIdInt))

	builder := flatbuffers.NewBuilder(1024)

	clubIdBuffer := builder.CreateString(clubId)
	event := builder.CreateString(eventToSend)

	messagePayload := CreateExitClubPayload(builder, userIds)

	msg := CreateEventWrapper(builder, clubIdBuffer, event, eventSchemaHandler.PayloadExitClubPayload, messagePayload)

	builder.Finish(msg)

	return builder.FinishedBytes()
}

func CreateExitClubPayload(builder *flatbuffers.Builder, userIds []int32) flatbuffers.UOffsetT {
	eventSchemaHandler.ExitClubPayloadStartUserIdsVector(builder, len(userIds))

	for _, userId := range userIds {
		// userId, _ := strconv.Atoi(micEnabledUserRef.Ref.ID)
		// userId := utils.StringToInt32(micEnabledUserRef.Ref.ID)
		builder.PrependInt32(userId)
	}
	userIdsPayload := builder.EndVector(len(userIds))
	eventSchemaHandler.ExitClubPayloadStart(builder)

	eventSchemaHandler.ExitClubPayloadAddUserIds(builder, userIdsPayload)

	return eventSchemaHandler.ExitClubPayloadEnd(builder)
}
