package builder

class Guide(private val builder: Builder) {
    fun doGuide() {
        builder.setTitle("testTitle")
        builder.addString("testString")
        builder.addItems(arrayOf("test1", "test2"))
        builder.close()
    }
}