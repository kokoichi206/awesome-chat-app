package data.remote.model

import domain.model.MessageType
import domain.model.RoomMessage
import kotlinx.datetime.TimeZone
import kotlinx.datetime.toInstant
import kotlinx.datetime.toLocalDateTime
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class RoomMessageDao(
    val id: String = "",
    @SerialName("user_id") val userId: String = "",
    val type: String = "",
    val content: String = "",
    val timestamp: String = "",
)

fun RoomMessageDao.toRoomMessage(): RoomMessage {
    return RoomMessage(
        id = id,
        userId = userId,
        type = when (type) {
            "text" -> MessageType.TEXT
            "image" -> MessageType.IMAGE
            "stamp" -> MessageType.STAMP
            else -> MessageType.UNKNOWN
        },
        content = content,
        time = timestamp.toInstant().toLocalDateTime(TimeZone.currentSystemDefault())
    )
}
