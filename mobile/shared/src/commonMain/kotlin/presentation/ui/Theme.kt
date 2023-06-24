package presentation.ui

import androidx.compose.material.MaterialTheme
import androidx.compose.material.lightColors
import androidx.compose.runtime.Composable
import androidx.compose.ui.graphics.Color

@Composable
fun CustomChatTheme(
    content: @Composable () -> Unit
) {
    MaterialTheme(
        colors = chatLightColors,
        typography = Typography,
        content = content
    )
}

private val chatLightColors = lightColors(
    primary = Color(0xFF79E278),
    // トークでの背景色。
    secondary = Color(0xFF8CABD8),
    background = Color.White,
    onBackground = Color.White,
)
