import UIKit
import SwiftUI
import shared

struct ComposeView: UIViewControllerRepresentable {
    func makeUIViewController(context: Context) -> UIViewController {
        Main_iosKt.MainViewController(idToken: "")
    }

    func updateUIViewController(_ uiViewController: UIViewController, context: Context) {}
}

struct ContentView: View {

    @State private var showSignInView: Bool = false

    var body: some View {
        ZStack {
            NavigationStack {
                ComposeView()
                    .ignoresSafeArea(.all)
                VStack {
                    // TODO: 固定のカラーコードの使用を止める。
                    Color(red: 140/255, green: 171/255, blue: 216/255, opacity: 1)
                        .ignoresSafeArea(edges: .top).frame(height: 0)
                    Spacer()
                    Color(red: 140/255, green: 171/255, blue: 216/255, opacity: 1)
                        .ignoresSafeArea(edges: .bottom).frame(height: 0)
                }
            }
        }
        .onAppear {
            let authUser = try? AuthenticationManager.shared.getAuthenticatedUser()
            self.showSignInView = authUser == nil
        }
        .fullScreenCover(isPresented: $showSignInView, content: {
            NavigationStack {
                AuthenticationView(showSignInView: $showSignInView)
            }
        })
        .preferredColorScheme(.dark)

    }
}
