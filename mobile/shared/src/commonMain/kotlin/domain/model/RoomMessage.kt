package domain.model

import kotlinx.datetime.LocalDateTime

data class RoomMessage(
    val id: String,
    val userId: String,
    val type: MessageType,
    val content: String,
    val time: LocalDateTime,
)

enum class MessageType(val type: String) {
    TEXT("text"),
    IMAGE("image"),
    STAMP("stamp"),
    UNKNOWN("unknown"),
}
