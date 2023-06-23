package presentation.chat

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Icon
import androidx.compose.material.MaterialTheme
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.Send
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import presentation.chat.components.CustomTextField
import presentation.chat.components.OneMessage
import presentation.ui.AppColors
import presentation.ui.IconSizeMedium
import presentation.ui.SpaceSmall

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

    Column(
        modifier = Modifier
            .fillMaxSize(),
    ) {
        LazyColumn(
            modifier = Modifier
                .weight(1f)
                .background(MaterialTheme.colors.secondary),
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

        SendArea(
            state = state,
            onValueChange = { viewModel.setInput(it) },
            onSendClick = { viewModel.onSendClick() },
        )
    }
}

@Composable
fun SendArea(
    state: ChatState,
    onValueChange: (String) -> Unit = {},
    onSendClick: () -> Unit = {},
) {
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .background(Color.White.copy(alpha = 0f))
            .padding(SpaceSmall),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Icon(
            modifier = Modifier
                .size(IconSizeMedium),
            imageVector = Icons.Default.KeyboardArrowRight,
            contentDescription = "arrow right",
            tint = AppColors.DarkGray,
        )

        CustomTextField(
            modifier = Modifier
                .weight(1f)
                .clip(RoundedCornerShape(16.dp))
                .background(AppColors.LightGray)
                .padding(horizontal = SpaceSmall),
            value = state.input,
            onValueChange = onValueChange,
            textStyle = MaterialTheme.typography.body1,
            singleLine = true,
        )

        Spacer(modifier = Modifier.width(SpaceSmall))

        Icon(
            modifier = Modifier
                .clickable {
                    onSendClick()
                },
            imageVector = Icons.Default.Send,
            contentDescription = "send icon",
            tint = if (state.input.isBlank()) AppColors.LightGray else AppColors.ButtonBlue,
        )
    }
}
