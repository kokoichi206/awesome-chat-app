package data.remote

import data.remote.model.PostMessageBody
import data.remote.model.RoomMessagesDao
import data.remote.model.UsersDto
import domain.model.RoomMessage
import io.ktor.client.call.body
import io.ktor.client.request.get
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.http.ContentType
import io.ktor.http.contentType
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

    suspend fun postMessage(roomMessage: RoomMessage, roomId: String): Unit = withContext(dispatcher.io) {
        client.post {
            pathUrl("rooms/$roomId/messages")
            contentType(ContentType.Application.Json)
            setBody(
                PostMessageBody(
                    userId = roomMessage.userId,
                    type = roomMessage.type.toString(),
                    content = roomMessage.content,
                )
            )
        }.body()
    }
}
