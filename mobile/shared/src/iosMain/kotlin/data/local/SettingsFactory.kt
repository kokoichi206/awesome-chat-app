package data.local

import com.russhwolf.settings.ExperimentalSettingsImplementation
import com.russhwolf.settings.KeychainSettings
import com.russhwolf.settings.Settings
import di.PlatformModule

actual class SettingsFactory actual constructor(platformModule: PlatformModule) {
    private val serviceName = "awesome-chat-app"

    @OptIn(ExperimentalSettingsImplementation::class)
    actual fun create(): Settings {
        return KeychainSettings(service = serviceName)
    }
}
