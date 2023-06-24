package domain.usecase

import domain.model.RoomMessage
import domain.repository.RoomRepository
import io.ktor.client.plugins.ClientRequestException
import io.ktor.utils.io.errors.IOException
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import util.Resource

class GetRoomMessagesUsecase(
    private val repository: RoomRepository,
) {
    fun execute(roomId: String): Flow<Resource<List<RoomMessage>>> = flow {
        try {
            emit(Resource.Loading())
            val messages = repository.getMessages(roomId)
            emit(Resource.Success(messages))
        } catch (e: IOException) {
            emit(Resource.Error(e.message ?: "Couldn't reach server. Check your network connection."))
        } catch (e: ClientRequestException) {
            emit(Resource.Error(e.message))
        }
    }
}
