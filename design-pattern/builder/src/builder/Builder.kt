package builder

interface Builder {
    fun setTitle(title: String)
    fun addString(str: String)
    fun addItems(items: Array<String>)
    fun close()
}