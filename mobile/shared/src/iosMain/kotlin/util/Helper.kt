package util

import di.sharedModules
import org.koin.core.context.startKoin

fun initKoin(){
    startKoin {
        modules(sharedModules())
    }
}
