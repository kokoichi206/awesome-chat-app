import androidx.compose.material.MaterialTheme
import androidx.compose.runtime.Composable
import presentation.chat.ChatScreen

@Composable
fun App() {
    // FIXME: navigation した時に、遷移元から値を受け取る。
    val roomId = "e8e31f8a-b0be-425e-8101-95e8c84bc699"

    MaterialTheme {
        ChatScreen(
            roomId = roomId,
        )
    }
}

expect fun getPlatformName(): String
