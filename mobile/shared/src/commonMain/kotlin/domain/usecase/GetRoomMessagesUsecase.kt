package domain.usecase

import domain.model.RoomMessage
import domain.repository.RoomRepository

class GetRoomMessagesUsecase(
    private val repository: RoomRepository,
) {
    suspend fun execute(roomId: String): List<RoomMessage> {
        return repository.getMessages(roomId)
    }
}