import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*


suspend fun main() {
    val client = HttpClient(CIO)
    val response: HttpResponse = client.get("https://ktor.io/")
//    val channel
    val sendMessage: HttpResponse = client.post("https://discord.com/api/channels/{channel.id}/messages")
    println(response.status)
    client.close()
}