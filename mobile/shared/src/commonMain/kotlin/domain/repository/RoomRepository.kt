package domain.repository

import domain.model.MessageType
import domain.model.RoomMessage
import domain.model.User

interface RoomRepository {

    suspend fun getUsers(roomId: String): List<User>

    suspend fun getMessages(roomId: String): List<RoomMessage>

    suspend fun postMessage(roomMessage: RoomMessage, roomId: String)
}
