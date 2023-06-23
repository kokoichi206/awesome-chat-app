package domain.usecase

import domain.model.RoomMessage
import domain.repository.RoomRepository

class PostMessagesUsecase(
    private val repository: RoomRepository,
) {
    suspend fun execute(roomMessage: RoomMessage, roomId: String) {
        return repository.postMessage(roomMessage, roomId)
    }
}
