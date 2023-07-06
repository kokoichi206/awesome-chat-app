package data.local

import com.russhwolf.settings.Settings
import di.PlatformModule

expect class SettingsFactory(platformModule: PlatformModule) {
    fun create(): Settings
}
