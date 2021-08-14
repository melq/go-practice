package builder

class UpperTextBuilder : Builder {
    val buffer = StringBuffer()
    override fun setTitle(title: String) {
        buffer.append("====================\n")
        buffer.append("『${title.toUpperCase()}』\n\n")
    }

    override fun addString(str: String) {
        buffer.append("■${str.toUpperCase()}\n")
    }

    override fun addItems(items: Array<String>) {
        for (item in items) {
            buffer.append(" ・${item.toUpperCase()}\n")
        }
        buffer.append("\n")
    }

    override fun close() {
        buffer.append("====================\n")
    }
}