package data.remote

import data.remote.model.RoomMessagesDao
import data.remote.model.UsersDto
import io.ktor.client.call.body
import io.ktor.client.request.get
import kotlinx.coroutines.withContext
import util.Dispatcher

class RoomRemote(
    private val dispatcher: Dispatcher,
) {
    suspend fun getUsers(roomId: String): UsersDto = withContext(dispatcher.io) {
        client.get {
            pathUrl("rooms/$roomId/users")
        }.body()
    }

    suspend fun getMessages(roomId: String): RoomMessagesDao = withContext(dispatcher.io) {
        client.get {
            pathUrl("rooms/$roomId/messages")
        }.body()
    }
}
