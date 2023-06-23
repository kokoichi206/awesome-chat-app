package presentation.chat

import com.adeo.kviewmodel.KViewModel
import domain.model.MessageType
import domain.model.RoomMessage
import domain.usecase.GetRoomMessagesUsecase
import domain.usecase.GetRoomUsersUsecase
import domain.usecase.PostMessagesUsecase
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import kotlinx.datetime.Clock
import kotlinx.datetime.TimeZone
import kotlinx.datetime.toLocalDateTime
import org.koin.core.component.KoinComponent
import org.koin.core.component.inject

class ChatViewModel(
    private val roomId: String,
) : KViewModel(), KoinComponent {

    private val usersUsecase: GetRoomUsersUsecase by inject()
    private val messagesUsecase: GetRoomMessagesUsecase by inject()
    private val postMessageUsecase: PostMessagesUsecase by inject()

    private val _state = MutableStateFlow(ChatState())
    val state = _state.asStateFlow()

    init {
        getUsers(roomId)
        getMessages(roomId)
    }

    private fun getUsers(roomId: String) {
        viewModelScope.launch {
            val users = usersUsecase.execute(roomId)
            _state.update {
                it.copy(users = users)
            }
        }
    }

    private fun getMessages(roomId: String) {
        viewModelScope.launch {
            val messages = messagesUsecase.execute(roomId)
            _state.update {
                it.copy(messages = messages)
            }
        }
    }

    fun setInput(input: String) {
        _state.update {
            it.copy(input = input)
        }
    }

    fun onSendClick() {
        val roomMessage = RoomMessage(
            userId = state.value.myId,
            // TODO: テキスト以外の送信。
            type = MessageType.TEXT,
            content = state.value.input,
            time = Clock.System.now().toLocalDateTime(TimeZone.currentSystemDefault()),
        )

        viewModelScope.launch {
            postMessageUsecase.execute(roomMessage, roomId)
        }

        _state.update {
            it.copy(
                messages = it.messages.plus(roomMessage),
                input = "",
            )
        }
    }
}
