package data.remote.model

import domain.model.RoomMessage
import kotlinx.serialization.Serializable

@Serializable
data class RoomMessagesDao(
    val messages: List<RoomMessageDao>,
)

fun RoomMessagesDao.toRoomMessages(): List<RoomMessage> {
    return messages.map {
        it.toRoomMessage()
    }
}
