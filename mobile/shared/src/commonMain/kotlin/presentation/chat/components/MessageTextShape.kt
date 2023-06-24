package presentation.chat.components

import androidx.compose.ui.geometry.Rect
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Outline
import androidx.compose.ui.graphics.Path
import androidx.compose.ui.graphics.Shape
import androidx.compose.ui.unit.Density
import androidx.compose.ui.unit.LayoutDirection
import kotlin.math.PI
import kotlin.math.cos
import kotlin.math.sin

class MessageTextShape(private val cornerRadius: Float, private val isMyText: Boolean) : Shape {

    override fun createOutline(size: Size, layoutDirection: LayoutDirection, density: Density): Outline {
        return Outline.Generic(
            path = if (isMyText) {
                drawMyMessagePath(size = size, cornerRadius = cornerRadius)
            } else {
                drawOtherMessagePath(size = size, cornerRadius = cornerRadius)
            }
        )
    }
}

fun drawMyMessagePath(size: Size, cornerRadius: Float): Path {
    return Path().apply {
        reset()

        // Top left arc
        arcTo(
            rect = Rect(
                left = 0f,
                top = 0f,
                right = 2 * cornerRadius,
                bottom = 2 * cornerRadius,
            ),
            startAngleDegrees = 180.0f,
            sweepAngleDegrees = 90.0f,
            forceMoveTo = false,
        )
        lineTo(x = size.width - cornerRadius, y = 0f)

        // Top right arc
        arcTo(
            rect = Rect(
                left = size.width - 2 * cornerRadius,
                top = 0f,
                right = size.width,
                bottom = 2 * cornerRadius,
            ),
            startAngleDegrees = 270.0f,
            sweepAngleDegrees = 30.0f,
            forceMoveTo = false,
        )
        lineTo(x = size.width, y = 0f)
        lineTo(
            x = size.width - cornerRadius * (1 - cos(PI / 6)).toFloat(),
            y = cornerRadius * (1 - sin(PI / 6)).toFloat()
        )
        arcTo(
            rect = Rect(
                left = size.width - 2 * cornerRadius,
                top = 0f,
                right = size.width,
                bottom = 2 * cornerRadius,
            ),
            startAngleDegrees = 330.0f,
            sweepAngleDegrees = 30.0f,
            forceMoveTo = false,
        )
        lineTo(x = size.width, y = size.height - cornerRadius)

        // Bottom right arc
        arcTo(
            rect = Rect(
                left = size.width - 2 * cornerRadius,
                top = size.height - 2 * cornerRadius,
                right = size.width,
                bottom = size.height,
            ),
            startAngleDegrees = 0.0f,
            sweepAngleDegrees = 90.0f,
            forceMoveTo = false,
        )
        lineTo(x = cornerRadius, y = size.height)

        // Bottom left arc
        arcTo(
            rect = Rect(
                left = 0.0f,
                top = size.height - 2 * cornerRadius,
                right = 2 * cornerRadius,
                bottom = size.height,
            ),
            startAngleDegrees = 90.0f,
            sweepAngleDegrees = 90.0f,
            forceMoveTo = false,
        )
        lineTo(x = 0f, y = cornerRadius)

        close()
    }
}

fun drawOtherMessagePath(size: Size, cornerRadius: Float): Path {
    return Path().apply {
        reset()

        // Top left arc
        arcTo(
            rect = Rect(
                left = 0f,
                top = 0f,
                right = 2 * cornerRadius,
                bottom = 2 * cornerRadius,
            ),
            startAngleDegrees = 180.0f,
            sweepAngleDegrees = 30.0f,
            forceMoveTo = false,
        )
        lineTo(x = 0f, y = 0f)
        lineTo(
            x = cornerRadius * (1 - sin(PI / 6)).toFloat(),
            y = cornerRadius * (1 - cos(PI / 6)).toFloat()
        )
        arcTo(
            rect = Rect(
                left = 0f,
                top = 0f,
                right = 2 * cornerRadius,
                bottom = 2 * cornerRadius,
            ),
            startAngleDegrees = 240.0f,
            sweepAngleDegrees = 30.0f,
            forceMoveTo = false,
        )

        // Top right arc
        arcTo(
            rect = Rect(
                left = size.width - 2 * cornerRadius,
                top = 0f,
                right = size.width,
                bottom = 2 * cornerRadius,
            ),
            startAngleDegrees = 270.0f,
            sweepAngleDegrees = 90.0f,
            forceMoveTo = false,
        )
        lineTo(x = size.width, y = size.height - cornerRadius)

        // Bottom right arc
        arcTo(
            rect = Rect(
                left = size.width - 2 * cornerRadius,
                top = size.height - 2 * cornerRadius,
                right = size.width,
                bottom = size.height,
            ),
            startAngleDegrees = 0.0f,
            sweepAngleDegrees = 90.0f,
            forceMoveTo = false,
        )
        lineTo(x = cornerRadius, y = size.height)

        // Bottom left arc
        arcTo(
            rect = Rect(
                left = 0.0f,
                top = size.height - 2 * cornerRadius,
                right = 2 * cornerRadius,
                bottom = size.height,
            ),
            startAngleDegrees = 90.0f,
            sweepAngleDegrees = 90.0f,
            forceMoveTo = false,
        )
        lineTo(x = 0f, y = cornerRadius)

        close()
    }
}
