import builder.Guide
import builder.TextBuilder
import builder.UpperTextBuilder

fun main(args: Array<String>) {
    val textBuilder = TextBuilder()
    Guide(textBuilder).doGuide()
    val upperTextBuilder = UpperTextBuilder()
    Guide(upperTextBuilder).doGuide()

    println(textBuilder.buffer.toString())
    println(upperTextBuilder.buffer.toString())
}