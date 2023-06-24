package domain.model

data class User(
    val id: String,
    val imgUrl: String?,
    val name: String,
    val lastReadAt: String,
)
