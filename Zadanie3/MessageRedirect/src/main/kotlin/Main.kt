
import dev.kord.core.Kord
import dev.kord.core.entity.ReactionEmoji
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json
import java.io.File
import java.util.*


suspend fun main() {
    val props = Properties().apply {
        load(File(".env").inputStream())
    }
    val webhookUrl = props.getProperty("webhookUrl")
    val discordBotToken = props.getProperty("discordBotToken")

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

    val kord = Kord(discordBotToken)
    kord.on<MessageCreateEvent> {
        if(message.author?.isBot == false){
            if(message.content.contains( "/categories")){
                message.channel.createMessage(Data.getAllCategories())
            }
            else{
                message.addReaction(ReactionEmoji.Unicode("❤️"))
                println("Message ${message.content}")
            }
        }
    }
    kord.login()

}
@Serializable
data class DiscordMessage(val content: String)