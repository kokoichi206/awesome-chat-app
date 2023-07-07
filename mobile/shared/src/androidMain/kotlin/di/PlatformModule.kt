package di

import AppContext
import android.app.Application

actual class PlatformModule constructor(
    private val application: Application,
) {
    actual val appContext: AppContext
        get() = application
}
