package jp.mydns.kokoichi206.awesomechatapp

import MainView
import android.os.Bundle
import android.widget.Toast
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.compose.setContent
import androidx.activity.result.IntentSenderRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.appcompat.app.AppCompatActivity
import androidx.compose.foundation.isSystemInDarkTheme
import androidx.compose.runtime.DisposableEffect
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.ui.graphics.Color
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.google.accompanist.systemuicontroller.rememberSystemUiController
import com.google.android.gms.auth.api.identity.Identity
import jp.mydns.kokoichi206.awesomechatapp.presentation.sign_in.GoogleAuthUiClient
import jp.mydns.kokoichi206.awesomechatapp.presentation.sign_in.SignInScreen
import jp.mydns.kokoichi206.awesomechatapp.presentation.sign_in.SignInViewModel
import kotlinx.coroutines.launch

class MainActivity : AppCompatActivity() {

    private val googleAuthUiClient by lazy {
        GoogleAuthUiClient(
            context = applicationContext,
            oneTapClient = Identity.getSignInClient(applicationContext),
        )
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        var idToken: String? = null

        setContent {
            val systemUiController = rememberSystemUiController()
            val useDarkIcons = !isSystemInDarkTheme()

            val navController = rememberNavController()

            NavHost(navController = navController, startDestination = "sign_in") {
                composable("sign_in") {
                    val viewModel = viewModel<SignInViewModel>()
                    val state by viewModel.state.collectAsStateWithLifecycle()

                    DisposableEffect(systemUiController, useDarkIcons) {
                        systemUiController.setSystemBarsColor(
                            color = Color.White,
                            darkIcons = useDarkIcons
                        )

                        onDispose {}
                    }

                    LaunchedEffect(key1 = Unit) {
                        googleAuthUiClient.getSignedInUser()?.let {
                            it.email?.let {
                                // TODO: navigate 条件は？
                                navController.navigate("chat")
                            }
                        }
                    }

                    val launcher = rememberLauncherForActivityResult(
                        contract = ActivityResultContracts.StartIntentSenderForResult(),
                        onResult = { result ->
                            if (result.resultCode == RESULT_OK) {
                                lifecycleScope.launch {
                                    val signInResult = googleAuthUiClient.signInWithIntent(
                                        intent = result.data ?: return@launch,
                                    )

                                    viewModel.onSignInResult(signInResult)
                                }
                            }
                        },
                    )

                    LaunchedEffect(key1 = state.isSignInSuccessful) {
                        if (state.isSignInSuccessful) {
                            Toast.makeText(
                                applicationContext,
                                "Sign in successful",
                                Toast.LENGTH_LONG,
                            ).show()

                            idToken = state.idToken
                            navController.navigate("chat")
                        }
                    }

                    SignInScreen(
                        state = state,
                        onSignInClick = {
                            lifecycleScope.launch {
                                val signInIntentSender = googleAuthUiClient.signIn()
                                launcher.launch(
                                    IntentSenderRequest.Builder(
                                        signInIntentSender ?: return@launch,
                                    ).build()
                                )
                            }
                        },
                    )
                }


                composable(
                    route = "chat",
                ) {
                    DisposableEffect(key1 = Unit) {
                        systemUiController.setSystemBarsColor(
                            color = Color(0xFF8CABD8),
                            darkIcons = useDarkIcons
                        )

                        onDispose {}
                    }

                    LaunchedEffect(true) {
                        if (idToken == null) {
                            lifecycleScope.launch {
                                googleAuthUiClient.signOut()
                            }
                        }
                    }

                    if (idToken == null) {
                        navController.navigate("sign_in")
                    } else {
                        MainView(idToken!!)
                    }
                }
            }
        }
    }
}
