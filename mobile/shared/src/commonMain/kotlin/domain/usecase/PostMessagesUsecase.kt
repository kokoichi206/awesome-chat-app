package domain.usecase

import domain.model.RoomMessage
import domain.repository.RoomRepository
import io.ktor.client.plugins.ClientRequestException
import io.ktor.utils.io.errors.IOException
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import util.Resource

class PostMessagesUsecase(
    private val repository: RoomRepository,
) {
    fun execute(roomMessage: RoomMessage, roomId: String): Flow<Resource<Unit>> = flow {
        try {
            emit(Resource.Loading())
            repository.postMessage(roomMessage, roomId)
            emit(Resource.Success(Unit))
        } catch (e: IOException) {
            emit(Resource.Error(e.message ?: "Couldn't reach server. Check your network connection."))
        } catch (e: ClientRequestException) {
            emit(Resource.Error(e.message))
        }
    }
}
