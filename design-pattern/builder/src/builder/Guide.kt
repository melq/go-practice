package builder

class Guide(private val builder: Builder) {
    fun doGuide() {
        builder.setTitle("TEST_TITLE")
        builder.addString("testString")
        builder.addItems(arrayOf("test1", "test2"))
        builder.close()
    }
}