import UIKit
import SwiftUI
import shared

struct ComposeView: UIViewControllerRepresentable {

    @Binding var idToken: String

    func makeUIViewController(context: Context) -> UIViewController {
        Main_iosKt.MainViewController(idToken: idToken)
    }

    func updateUIViewController(_ uiViewController: UIViewController, context: Context) {}
}

struct ContentView: View {

    @State private var showSignInView: Bool = true
    @State private var idToken: String = ""

    var body: some View {
        ZStack {
            NavigationStack {
                if self.idToken != "" {
                    ComposeView(idToken: $idToken)
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
        }
        .onAppear {
            let authUser = try? AuthenticationManager.shared.getAuthenticatedUser()
            self.showSignInView = authUser == nil
        }
        .fullScreenCover(isPresented: $showSignInView, content: {
            NavigationStack {
                AuthenticationView(showSignInView: $showSignInView, idToken: $idToken)
            }
        })
        .preferredColorScheme(.dark)

    }
}
