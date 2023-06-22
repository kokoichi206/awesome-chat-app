import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material.BottomAppBar
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Scaffold
import androidx.compose.material.Text
import androidx.compose.material.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import presentation.chat.ChatScreen

@Composable
fun App() {
    // FIXME: navigation した時に、遷移元から値を受け取る。
    val roomId = "e8e31f8a-b0be-425e-8101-95e8c84bc699"
    val roomName = "部屋の名前"

    MaterialTheme {
        ChatScreen(
            roomId = roomId,
        )
        Scaffold(
            modifier = Modifier
                .fillMaxSize(),
            topBar = {
                TopAppBar(
                    title = {
                        Text(">  $roomName")
                    },
                    backgroundColor = Color(0xFF8CABD8),
                )
            },
        ) {
            Box(
                modifier = Modifier
                    .fillMaxSize()
                    .padding(it)
            ) {
                ChatScreen(
                    roomId = roomId,
                )
            }
        }
    }
}

expect fun getPlatformName(): String
