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

    func signInGoogle() async throws -> String {

        let helper = SignInGoogleHelper()
        let token = try await helper.signIn()
        try await AuthenticationManager.shared.signInWithGoogle(token: token)
        return token.idToken
    }
}

struct AuthenticationView: View {

    @StateObject private var viewModel = AuthenticationViweModel()
    @Binding var showSignInView: Bool
    @Binding var idToken: String

    var body: some View {
        VStack {
            GoogleSignInButton(viewModel: GoogleSignInButtonViewModel(scheme: .dark, style: .wide, state: .normal)) {
                Task {
                    do {
                        idToken = try await viewModel.signInGoogle()
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
    @State private static var idToken = ""

    static var previews: some View {
        if #available(iOS 16.0, *) {
            NavigationStack {
                AuthenticationView(showSignInView: .constant(false), idToken: $idToken)
            }
        } else {
            // Fallback on earlier versions
        }
    }
}
