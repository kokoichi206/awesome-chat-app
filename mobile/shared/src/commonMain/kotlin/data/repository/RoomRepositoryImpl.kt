package data.repository

import data.remote.RoomRemote
import data.remote.model.toRoomMessages
import data.remote.model.toUsers
import domain.model.RoomMessage
import domain.model.User
import domain.repository.RoomRepository

class RoomRepositoryImpl(
    private val remote: RoomRemote,
) : RoomRepository {

    override suspend fun getUsers(roomId: String): List<User> {
        return remote.getUsers(roomId).toUsers()
    }

    override suspend fun getMessages(roomId: String): List<RoomMessage> {
        return remote.getMessages(roomId).toRoomMessages()
    }
}