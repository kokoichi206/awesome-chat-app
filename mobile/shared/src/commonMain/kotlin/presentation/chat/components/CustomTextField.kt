package presentation.chat.components

import androidx.compose.foundation.background
import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.foundation.text.BasicTextField
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.ExperimentalMaterialApi
import androidx.compose.material.TextFieldDefaults
import androidx.compose.material.TextFieldDefaults.indicatorLine
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.input.VisualTransformation
import androidx.compose.ui.unit.dp

@OptIn(ExperimentalMaterialApi::class)
@Composable
fun CustomTextField(
    modifier: Modifier = Modifier,
    value: String,
    enabled: Boolean = true,
    singleLine: Boolean = true,
    visualTransformation: VisualTransformation = VisualTransformation.None,
    interactionSource: MutableInteractionSource = remember { MutableInteractionSource() },
    isError: Boolean = false,
    keyboardOptions: KeyboardOptions = KeyboardOptions.Default,
    keyboardActions: KeyboardActions = KeyboardActions.Default,
    onValueChange: (String) -> Unit,
    textStyle: TextStyle = TextStyle.Default,
) {
    BasicTextField(
        value = value,
        modifier = modifier
            .background(Color.Transparent, TextFieldDefaults.TextFieldShape)
            .indicatorLine(
                enabled,
                isError,
                interactionSource,
                TextFieldDefaults.textFieldColors(
                    focusedIndicatorColor = Color.Transparent,
                )
            ),
        onValueChange = { onValueChange(it) },
        enabled = enabled,
        keyboardOptions = keyboardOptions,
        keyboardActions = keyboardActions,
        interactionSource = interactionSource,
        singleLine = singleLine,
        textStyle = textStyle,
    ) { innerTextField ->
        TextFieldDefaults.TextFieldDecorationBox(
            value = value,
            innerTextField = innerTextField,
            enabled = enabled,
            singleLine = singleLine,
            visualTransformation = visualTransformation,
            interactionSource = interactionSource,

            contentPadding = TextFieldDefaults.textFieldWithLabelPadding(0.dp),
        )
    }
}
