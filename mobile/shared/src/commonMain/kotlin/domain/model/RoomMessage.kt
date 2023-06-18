package domain.model

data class RoomMessage(
    val id: String,
    val userId: String,
    val type: String,
    val content: String,
    val timestamp: String,
)
