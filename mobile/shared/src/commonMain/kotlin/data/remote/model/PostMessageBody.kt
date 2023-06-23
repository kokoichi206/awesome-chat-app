package data.remote.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class PostMessageBody(
    @SerialName("user_id") val userId: String,
    val type: String,
    val content: String,
)
