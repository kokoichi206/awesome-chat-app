package data.remote

import io.ktor.client.HttpClient
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.HttpRequestBuilder
import io.ktor.http.appendPathSegments
import io.ktor.http.takeFrom
import io.ktor.serialization.kotlinx.json.json
import kotlinx.serialization.json.Json
import util.BuildKonfig

internal val client = HttpClient {
    install(ContentNegotiation) {
        json(Json {
            ignoreUnknownKeys = true
            useAlternativeNames = false
            isLenient = true
        })
    }

    expectSuccess = true
}

internal fun HttpRequestBuilder.pathUrl(path: String) {
    url {
        takeFrom(BuildKonfig.BASE_URL)
        appendPathSegments(BuildKonfig.API_PATH)
        appendPathSegments(path)
    }
}
