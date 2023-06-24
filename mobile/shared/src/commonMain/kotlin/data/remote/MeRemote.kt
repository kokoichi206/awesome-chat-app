package data.remote

import data.remote.model.UserDto
import io.ktor.client.call.body
import io.ktor.client.request.get
import kotlinx.coroutines.withContext
import util.Dispatcher

class MeRemote(
    private val dispatcher: Dispatcher,
) {
    suspend fun getMe(): UserDto = withContext(dispatcher.io) {
        client.get {
            pathUrl("users/me")
        }.body()
    }
}
