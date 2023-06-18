package data.remote

import io.ktor.client.HttpClient
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.HttpRequestBuilder
import io.ktor.http.appendPathSegments
import io.ktor.http.takeFrom
import io.ktor.serialization.kotlinx.json.json
import kotlinx.serialization.json.Json

private const val BASE_URL = "http://192.168.0.113:8383"
private const val API_PATH = "api"

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
        takeFrom(BASE_URL)
        appendPathSegments(API_PATH)
        appendPathSegments(path)
    }
}
