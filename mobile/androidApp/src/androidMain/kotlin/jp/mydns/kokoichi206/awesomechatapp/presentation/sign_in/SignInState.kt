package jp.mydns.kokoichi206.awesomechatapp.presentation.sign_in

data class SignInState(
    val isSignInSuccessful: Boolean = false,
    val signInError: String? = null,
    val email: String? = null,
    val idToken: String? = null,
)
