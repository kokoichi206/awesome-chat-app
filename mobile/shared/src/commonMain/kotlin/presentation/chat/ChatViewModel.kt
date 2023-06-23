package presentation.chat

import com.adeo.kviewmodel.KViewModel
import domain.model.MessageType
import domain.model.RoomMessage
import domain.usecase.GetRoomMessagesUsecase
import domain.usecase.GetRoomUsersUsecase
import domain.usecase.PostMessagesUsecase
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.launchIn
import kotlinx.coroutines.flow.onEach
import kotlinx.coroutines.flow.update
import kotlinx.datetime.Clock
import kotlinx.datetime.TimeZone
import kotlinx.datetime.toLocalDateTime
import org.koin.core.component.KoinComponent
import org.koin.core.component.inject
import util.Resource

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
        usersUsecase.execute(roomId).onEach { result ->
            when (result) {
                is Resource.Success -> {
                    _state.update {
                        it.copy(
                            isLoading = false,
                            users = result.data!!,
                        )
                    }
                }

                is Resource.Loading -> {
                    _state.update { it.copy(isLoading = true) }
                }

                is Resource.Error -> {
                    _state.update { it.copy(isLoading = false, errorMessage = result.message) }
                }
            }
        }.launchIn(viewModelScope)
    }

    private fun getMessages(roomId: String) {
        messagesUsecase.execute(roomId).onEach { result ->
            when (result) {
                is Resource.Success -> {
                    _state.update {
                        it.copy(
                            isLoading = false,
                            messages = result.data!!,
                        )
                    }
                }

                is Resource.Loading -> {
                    _state.update { it.copy(isLoading = true) }
                }

                is Resource.Error -> {
                    _state.update { it.copy(isLoading = false, errorMessage = result.message) }
                }
            }
        }.launchIn(viewModelScope)
    }

    fun setInput(input: String) {
        _state.update {
            it.copy(input = input)
        }
    }

    fun onSendClick() {
        if (state.value.input.isBlank()) {
            return
        }

        val roomMessage = RoomMessage(
            userId = state.value.myId,
            // TODO: テキスト以外の送信。
            type = MessageType.TEXT,
            content = state.value.input,
            time = Clock.System.now().toLocalDateTime(TimeZone.currentSystemDefault()),
        )

        postMessageUsecase.execute(roomMessage, roomId).onEach { result ->
            when (result) {
                is Resource.Success -> {
                    _state.update { it.copy(isLoading = false) }
                }

                is Resource.Loading -> {
                    _state.update { it.copy(isLoading = true) }
                }

                is Resource.Error -> {
                    _state.update { it.copy(isLoading = false, errorMessage = result.message) }
                }
            }
        }.launchIn(viewModelScope)

        _state.update {
            it.copy(
                messages = it.messages.plus(roomMessage),
                input = "",
            )
        }
    }
}
