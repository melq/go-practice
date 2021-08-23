import java.util.*

data class Card(val pattern: String, val number: Int)

fun swap(list: MutableList<Card>, i: Int, j: Int) {
    val tmp = list[i]
    list[i] = list[j]
    list[j] = tmp
}

fun partition(list: MutableList<Card>, p: Int, r: Int): Int {
    val x = list[r].number
    var i = p - 1
    for (j in p until r) {
        if (list[j].number <= x) {
            i++
            swap(list, i, j)
        }
    }
    swap(list, i + 1, r)
    return i + 1
}

fun quickSort(list: MutableList<Card>, p: Int, r: Int) {
    if (p < r) {
        val q = partition(list, p, r)
        quickSort(list, p, q - 1)
        quickSort(list, q + 1, r)
    }
}

fun main(args: Array<String>) {
    val scanner = Scanner(System.`in`)
    val n = scanner.nextInt()
    val list = mutableListOf<Card>()
    for (i in 0 until n) {
        list.add(Card(
            scanner.next(),
            scanner.nextInt())
        )
    }

    quickSort(list, 0, n - 1)
    for (card in list) {
        println("${card.pattern} ${card.number}")
    }
}