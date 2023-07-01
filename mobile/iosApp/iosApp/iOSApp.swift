import SwiftUI
import shared

import FirebaseCore
import FirebaseAuth
import GoogleSignIn

@main
struct iOSApp: App {
    init() {
        HelperKt.doInitKoin()

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
