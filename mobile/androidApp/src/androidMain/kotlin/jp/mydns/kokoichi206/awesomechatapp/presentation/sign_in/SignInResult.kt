package jp.mydns.kokoichi206.awesomechatapp.presentation.sign_in

data class SignInResult(
    val data: UserData?,
    val errorMessage: String?
)

data class UserData(
    val userId: String,
    val username: String?,
    val email: String?,
    val idToken: String? = null,
    val profilePictureUrl: String?,
)
