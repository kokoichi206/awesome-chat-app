package presentation.chat

import com.adeo.kviewmodel.KViewModel
import domain.usecase.GetRoomMessagesUsecase
import domain.usecase.GetRoomUsersUsecase
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import org.koin.core.component.KoinComponent
import org.koin.core.component.inject

class ChatViewModel(
    roomId: String,
) : KViewModel(), KoinComponent {

    private val usersUsecase: GetRoomUsersUsecase by inject()
    private val messagesUsecase: GetRoomMessagesUsecase by inject()

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
}
