package builder

class TextBuilder : Builder {
    val buffer: StringBuffer = StringBuffer()
    override fun setTitle(title: String) {
        buffer.append("====================\n")
        buffer.append("『$title』\n\n")
    }

    override fun addString(str: String) {
        buffer.append("■$str\n")
    }

    override fun addItems(items: Array<String>) {
        for (item in items) {
            buffer.append(" ・$item\n")
        }
        buffer.append("\n")
    }

    override fun close() {
        buffer.append("====================\n")
    }
}