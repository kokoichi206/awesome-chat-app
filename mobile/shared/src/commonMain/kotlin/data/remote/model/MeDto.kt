package data.remote.model

import domain.model.Me
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class MeDto(
    val email: String = "",
    val id: String = "",
    @SerialName("img_url") val imgUrl: String? = null,
    @SerialName("username") val userName: String = "",
)

fun MeDto.toMe(): Me {
    return Me(
        email = email,
        id = id,
        name = userName,
        imgUrl = imgUrl,
    )
}
