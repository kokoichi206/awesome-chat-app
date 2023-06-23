package domain.usecase

import domain.model.User
import domain.repository.RoomRepository
import io.ktor.client.plugins.ClientRequestException
import io.ktor.utils.io.errors.IOException
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import util.Resource

class GetRoomUsersUsecase(
    private val repository: RoomRepository,
) {
    fun execute(roomId: String): Flow<Resource<List<User>>> = flow {
        try {
            emit(Resource.Loading())
            val users = repository.getUsers(roomId)
            emit(Resource.Success(users))
        } catch (e: IOException) {
            emit(Resource.Error(e.message ?: "Couldn't reach server. Check your network connection."))
        } catch (e: ClientRequestException) {
            emit(Resource.Error(e.message))
        }
    }
}
