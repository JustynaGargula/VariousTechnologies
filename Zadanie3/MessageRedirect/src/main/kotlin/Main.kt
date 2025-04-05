import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json
import java.io.File
import java.util.Properties


suspend fun main() {
    val props = Properties().apply {
        load(File(".env").inputStream())
    }
    val webhookUrl = props.getProperty("webhookUrl")
    val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json(Json { prettyPrint = true })
        }
    }

    val response = client.post(webhookUrl) {
        contentType(ContentType.Application.Json)
        setBody(DiscordMessage("Test message"))
    }

    println("Status: ${response.status}")
    client.close()
}
@Serializable
data class DiscordMessage(val content: String)