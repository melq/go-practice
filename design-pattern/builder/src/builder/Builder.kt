package builder

interface Builder {
    fun setTitle(title: String)
    fun addString(str: String)
    fun addItem(item: Array<String>)
    fun close()
}