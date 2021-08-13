import builder.Guide
import builder.TextBuilder

fun main(args: Array<String>) {
    val builder = TextBuilder()
    Guide(builder).doGuide()
    println(builder.buffer.toString())
}