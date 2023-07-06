package data.local

import android.content.SharedPreferences
import androidx.security.crypto.EncryptedSharedPreferences
import androidx.security.crypto.MasterKey
import com.russhwolf.settings.Settings
import com.russhwolf.settings.SharedPreferencesSettings
import di.PlatformModule

actual class SettingsFactory actual constructor(private val platformModule: PlatformModule) {
    private val fileName = "awesome-chat-app-multiplatform-settings"

    actual fun create(): Settings {
        val masterKey = MasterKey.Builder(platformModule.appContext, MasterKey.DEFAULT_MASTER_KEY_ALIAS)
            .setKeyScheme(MasterKey.KeyScheme.AES256_GCM)
            .build()

        val prefs: SharedPreferences = EncryptedSharedPreferences.create(
            platformModule.appContext,
            fileName,
            masterKey,
            EncryptedSharedPreferences.PrefKeyEncryptionScheme.AES256_SIV,
            EncryptedSharedPreferences.PrefValueEncryptionScheme.AES256_GCM,
        )

        return SharedPreferencesSettings(prefs)
    }
}
