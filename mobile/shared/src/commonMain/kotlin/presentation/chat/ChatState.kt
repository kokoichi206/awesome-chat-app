package presentation.chat

import domain.model.RoomMessage
import domain.model.User

data class ChatState(
    // FIXME: api などで取得する。
    val myId: String = "9817faf8-5e8e-408e-a50d-972b82bc812d",
    val users: List<User> = emptyList(),
    val messages: List<RoomMessage> = emptyList(),
    val input: String = "",
)
