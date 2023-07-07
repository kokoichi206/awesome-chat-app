package util

import di.PlatformModule
import di.sharedModules
import org.koin.core.context.startKoin

fun initKoin(platformModule: PlatformModule){
    startKoin {
        modules(sharedModules(platformModule))
    }
}
