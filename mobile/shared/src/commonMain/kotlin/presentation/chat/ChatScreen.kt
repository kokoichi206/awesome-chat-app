package presentation.chat

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import presentation.chat.components.OneMessage

@Composable
fun ChatScreen(
    roomId: String,
) {
    val viewModel = ChatViewModel(roomId)

    ChatMainScreen(viewModel)
}

@Composable
fun ChatMainScreen(
    viewModel: ChatViewModel,
) {
    val state by viewModel.state.collectAsState()

    LazyColumn(
        modifier = Modifier
            .fillMaxSize()
            .background(Color(0xFF8CABD8)),
    ) {
        items(state.messages) { msg ->
            state.users.firstOrNull { user ->
                user.id == msg.userId
            }?.let {
                OneMessage(
                    myUid = state.myId,
                    roomMessage = msg,
                    user = it,
                )
            }
        }
    }
}
