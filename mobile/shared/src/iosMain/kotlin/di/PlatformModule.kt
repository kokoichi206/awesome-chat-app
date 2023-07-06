package di

import AppContext
import platform.darwin.NSObject

actual class PlatformModule {
    actual val appContext: AppContext
        get() = NSObject()
}
