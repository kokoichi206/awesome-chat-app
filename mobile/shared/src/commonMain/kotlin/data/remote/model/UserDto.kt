package data.remote.model

import domain.model.User
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class UserDto(
    val id: String = "",
    @SerialName("img_url") val imgUrl: String? = null,
    @SerialName("username") val userName: String = "",
    @SerialName("last_read_at") val lastReadAt: String = "",
)

fun UserDto.toUser(): User {
    return User(
        id = id,
        imgUrl = imgUrl,
        name = userName,
        lastReadAt = lastReadAt,
    )
}
