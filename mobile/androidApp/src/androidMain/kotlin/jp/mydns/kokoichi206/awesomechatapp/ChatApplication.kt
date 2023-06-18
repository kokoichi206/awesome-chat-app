package jp.mydns.kokoichi206.awesomechatapp

import android.app.Application
import di.sharedModules
import org.koin.core.context.startKoin

class ChatApplication: Application() {

    override fun onCreate() {
        super.onCreate()

        startKoin {
            modules(sharedModules())
        }
    }
}