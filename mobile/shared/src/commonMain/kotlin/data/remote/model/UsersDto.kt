package data.remote.model

import domain.model.User
import kotlinx.serialization.Serializable

@Serializable
data class UsersDto(
    val users: List<UserDto>,
)

fun UsersDto.toUsers(): List<User> {
    return users.map {
        it.toUser()
    }
}
