package presentation.chat.components

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.BoxWithConstraints
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.sizeIn
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.wrapContentSize
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.MaterialTheme
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.unit.dp
import dev.icerock.moko.resources.compose.painterResource
import dev.icerock.moko.resources.compose.stringResource
import domain.model.MessageType
import domain.model.RoomMessage
import domain.model.User
import io.kamel.core.Resource
import io.kamel.core.utils.cacheControl
import io.kamel.image.KamelImage
import io.kamel.image.asyncPainterResource
import io.ktor.client.utils.CacheControl
import jp.mydns.kokoichi206.awesomechatapp.resources.SharedRes
import presentation.ui.IconSizeLarge
import presentation.ui.SpaceMedium
import presentation.ui.SpaceSmall
import presentation.ui.SpaceTiny
import util.Constants

@Composable
fun OneMessage(
    myUid: String,
    roomMessage: RoomMessage,
    user: User,
) {
    val isMyText = myUid == roomMessage.userId
    Box(
        modifier = Modifier.fillMaxWidth(),
        contentAlignment = if (isMyText) Alignment.CenterEnd else Alignment.CenterStart,
    ) {
        if (isMyText) {
            MyMessage(roomMessage)
        } else {
            OthersMessage(roomMessage, user)
        }
    }
}

@Composable
fun MyMessage(
    roomMessage: RoomMessage,
) {
    Box(
        contentAlignment = Alignment.CenterEnd,
    ) {
        Row(
            modifier = Modifier
                .padding(SpaceTiny),
            verticalAlignment = Alignment.Bottom,
        ) {
            Column(
                modifier = Modifier
                    .padding(SpaceTiny),
                horizontalAlignment = Alignment.CenterHorizontally,
            ) {
                Text(
                    // FIXME: 他ユーザーからの既読情報取得。
                    text = stringResource(SharedRes.strings.read_message),
                    style = MaterialTheme.typography.body2,
                )

                val time = roomMessage.time
                Text(
                    text = "${time.hour}:${time.minute}",
                    style = MaterialTheme.typography.body2,
                )
            }

            when (roomMessage.type) {
                MessageType.TEXT -> MessageText(roomMessage, true)
                MessageType.IMAGE -> MessageImage(roomMessage)
                MessageType.STAMP -> TODO()
                MessageType.UNKNOWN -> TODO()
            }
        }
    }
}

@Composable
fun OthersMessage(
    roomMessage: RoomMessage,
    user: User,
) {
    BoxWithConstraints(
        contentAlignment = Alignment.CenterEnd,
    ) {
        Row(
            modifier = Modifier
                .padding(SpaceTiny),
            verticalAlignment = Alignment.Top,
        ) {
            Image(
                modifier = Modifier
                    .size(IconSizeLarge)
                    .clip(CircleShape)
                    .border(
                        shape = CircleShape,
                        width = 2.dp,
                        color = Color.Gray,
                    )
                    .padding(SpaceTiny),
                painter = painterResource(SharedRes.images.userdefault),
                contentDescription = "image of ${user.name}",
            )

            Row(
                modifier = Modifier
                    .padding(SpaceTiny),
                verticalAlignment = Alignment.Bottom,
            ) {
                when (roomMessage.type) {
                    MessageType.TEXT -> MessageText(roomMessage, false)
                    MessageType.IMAGE -> MessageImage(roomMessage)
                    MessageType.STAMP -> TODO()
                    MessageType.UNKNOWN -> TODO()
                }

                val time = roomMessage.time
                Text(
                    modifier = Modifier
                        .padding(SpaceTiny)
                        .width(32.dp),
                    text = stringResource(SharedRes.strings.post_time, time.hour, time.minute),
                    style = MaterialTheme.typography.body2,
                )
            }
        }
    }
}

@Composable
fun MessageText(
    roomMessage: RoomMessage,
    isMyText: Boolean,
) {
    Text(
        modifier = Modifier
            .sizeIn(maxWidth = Constants.MAX_MESSAGE_WIDTH)
            .wrapContentSize()
            .graphicsLayer {
                shadowElevation = 4.dp.toPx()
                shape = MessageTextShape(12.dp.toPx(), isMyText)
                clip = true
            }
            .background(color = if (isMyText) MaterialTheme.colors.primary else MaterialTheme.colors.onBackground)
            .padding(horizontal = SpaceMedium, vertical = SpaceSmall),
        text = roomMessage.content,
    )
}

@Composable
fun MessageImage(
    roomMessage: RoomMessage,
) {
    val painterResource: Resource<Painter> =
        asyncPainterResource(roomMessage.content) {
            requestBuilder {
                cacheControl(CacheControl.MAX_AGE)
            }
        }

    KamelImage(
        modifier = Modifier
            .sizeIn(
                maxHeight = Constants.MAX_MESSAGE_IMAGE_HEIGHT,
                maxWidth = Constants.MAX_MESSAGE_WIDTH,
            ),
        contentScale = ContentScale.Crop,
        resource = painterResource,
        contentDescription = "image of ${roomMessage.id}",
        onLoading = { progress -> CircularProgressIndicator(progress) },
        onFailure = { exception ->
            println(exception.message.toString())
        }
    )
}
