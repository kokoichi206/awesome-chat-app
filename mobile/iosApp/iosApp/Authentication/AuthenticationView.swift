//
//  AuthenticationView.swift
//  iosApp
//
//  Created by Takahiro Tominaga on 2023/07/01.
//  Copyright © 2023 orgName. All rights reserved.
//

import SwiftUI
import GoogleSignIn
import GoogleSignInSwift

@MainActor
final class AuthenticationViweModel: ObservableObject {

    func signInGoogle() async throws {

        let helper = SignInGoogleHelper()
        let token = try await helper.signIn()
        try await AuthenticationManager.shared.signInWithGoogle(token: token)
    }
}

struct AuthenticationView: View {

    @StateObject private var viewModel = AuthenticationViweModel()
    @Binding var showSignInView: Bool

    var body: some View {
        VStack {
            GoogleSignInButton(viewModel: GoogleSignInButtonViewModel(scheme: .dark, style: .wide, state: .normal)) {
                Task {
                    do {
                        try await viewModel.signInGoogle()
                        showSignInView = false
                    } catch {
                        print("ERROR \(error)")
                    }
                }
            }

            Spacer()
        }
        .padding()
        .navigationTitle("Sign in")
    }
}

struct AuthenticationView_Previews: PreviewProvider {
    static var previews: some View {
        if #available(iOS 16.0, *) {
            NavigationStack {
                AuthenticationView(showSignInView: .constant(false))
            }
        } else {
            // Fallback on earlier versions
        }
    }
}
