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
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import dev.icerock.moko.resources.compose.painterResource
import domain.model.RoomMessage
import domain.model.User
import jp.mydns.kokoichi206.awesomechatapp.resources.MR

@Composable
fun OneMessage(
    myUid: String,
    roomMessage: RoomMessage,
    user: User,
) {
    MessageText(roomMessage, user, myUid == roomMessage.userId)
}

@Composable
fun MessageText(
    roomMessage: RoomMessage,
    user: User,
    isMyText: Boolean,
) {
    if (isMyText) {
        Box(
            modifier = Modifier.fillMaxWidth(),
            contentAlignment = Alignment.CenterEnd,
        ) {
            MyMessageText(roomMessage)
        }
    } else {
        Box(
            modifier = Modifier.fillMaxWidth(),
            contentAlignment = Alignment.CenterStart,
        ) {
            OthersText(roomMessage, user)
        }
    }
}

@Composable
fun MyMessageText(
    roomMessage: RoomMessage,
) {
    Box(
        contentAlignment = Alignment.CenterEnd,
    ) {
        Row(
            modifier = Modifier
                .padding(4.dp),
        ) {
            Column(
                modifier = Modifier
                    .padding(4.dp),
                horizontalAlignment = Alignment.CenterHorizontally,
            ) {
                Text(
                    // FIXME: 他ユーザーからの既読情報取得。
                    text = "既読",
                    fontSize = 11.sp,
                )

                val time = roomMessage.time
                Text(
                    text = "${time.hour}:${time.minute}",
                    fontSize = 11.sp,
                )
            }

            Text(
                modifier = Modifier
                    .sizeIn(maxWidth = 240.dp)
                    .wrapContentSize()
                    .graphicsLayer {
                        shadowElevation = 4.dp.toPx()
                        shape = MessageTextShape(12.dp.toPx(), true)
                        clip = true
                    }
                    .background(color = Color(0xFF79E278))
                    .padding(horizontal = 16.dp, vertical = 8.dp),
                text = roomMessage.content,
            )
        }
    }
}

@Composable
fun OthersText(
    roomMessage: RoomMessage,
    user: User,
) {
    BoxWithConstraints(
        contentAlignment = Alignment.CenterEnd,
    ) {
        Row(
            modifier = Modifier
                .padding(4.dp),
            verticalAlignment = Alignment.Top,
        ) {
            Image(
                modifier = Modifier
                    .size(46.dp)
                    .clip(CircleShape)
                    .border(
                        shape = CircleShape,
                        width = 2.dp,
                        color = Color.Gray,
                    )
                    .padding(2.dp),
                painter = painterResource(MR.images.userdefault),
                contentDescription = "image of ${user.name}",
            )

            Row(
                modifier = Modifier
                    .padding(4.dp),
                verticalAlignment = Alignment.Bottom,
            ) {
                Text(
                    modifier = Modifier
                        .sizeIn(maxWidth = 240.dp)
                        .wrapContentSize()
                        .graphicsLayer {
                            shadowElevation = 4.dp.toPx()
                            shape = MessageTextShape(12.dp.toPx(), false)
                            clip = true
                        }
                        .background(color = Color.White)
                        .padding(horizontal = 16.dp, vertical = 8.dp),
                    text = roomMessage.content,
                )

                val time = roomMessage.time
                Text(
                    modifier = Modifier
                        .padding(4.dp)
                        .width(32.dp),
                    text = "${time.hour}:${time.minute}",
                    fontSize = 11.sp,
                )
            }
        }
    }
}
