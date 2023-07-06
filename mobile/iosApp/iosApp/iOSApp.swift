import SwiftUI
import shared

import FirebaseCore
import FirebaseAuth
import GoogleSignIn

@main
struct iOSApp: App {
    init() {
        let platformModule = PlatformModule()

        HelperKt.doInitKoin(platformModule: platformModule)

        setupAuthentication()
    }

    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}

extension iOSApp {
    private func setupAuthentication() {
        FirebaseApp.configure()
    }
}
