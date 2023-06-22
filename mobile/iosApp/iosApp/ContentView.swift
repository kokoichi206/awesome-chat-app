import UIKit
import SwiftUI
import shared

struct ComposeView: UIViewControllerRepresentable {
    func makeUIViewController(context: Context) -> UIViewController {
        Main_iosKt.MainViewController()
    }

    func updateUIViewController(_ uiViewController: UIViewController, context: Context) {}
}

struct ContentView: View {
    var body: some View {
        ZStack {
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
        }.preferredColorScheme(.dark)
    }
}
