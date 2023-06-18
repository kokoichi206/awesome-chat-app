package domain.usecase

import domain.model.User
import domain.repository.RoomRepository

class GetRoomUsersUsecase(
    private val repository: RoomRepository,
) {
    suspend fun execute(roomId: String): List<User> {
        return repository.getUsers(roomId)
    }
}