package data.remote.model

import domain.model.RoomMessage
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
        type = type,
        content = content,
        timestamp = timestamp,
    )
}
