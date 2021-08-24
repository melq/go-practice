import java.util.*

data class Card(val pattern: String, val number: Int)

fun swap(cards: MutableList<Card>, i: Int, j: Int) {
    val tmp = cards[i]
    cards[i] = cards[j]
    cards[j] = tmp
}

fun partition(cards: MutableList<Card>, p: Int, r: Int): Int {
    val x = cards[r].number
    var i = p - 1
    for (j in p until r) {
        if (cards[j].number <= x) {
            i++
            swap(cards, i, j)
        }
    }
    swap(cards, i + 1, r)
    return i + 1
}

fun quickSort(cards: MutableList<Card>, p: Int, r: Int) {
    if (p < r) {
        val q = partition(cards, p, r)
        quickSort(cards, p, q - 1)
        quickSort(cards, q + 1, r)
    }
}

fun isStable(originCards: MutableList<Card>, cards: MutableList<Card>): Boolean {
    var prev: Card
    val list1 = mutableListOf(cards[0])
    val list2 = mutableListOf<Card>()

    var i = 1
    while (i < cards.size) {
        prev = cards[i - 1]
        if (prev.number == cards[i].number) {
            list1.add(cards[i])
        } else {
            if (list1.size >= 2) {
                for (card in originCards) {
                    if (prev.number == card.number) {
                        list2.add(card)
                    }
                }
                if (list1 != list2) return false
                list2.clear()
            }
            list1.clear()
            list1.add(cards[i])
        }
        i++
    }
    return true
}

fun main(args: Array<String>) {
    val scanner = Scanner(System.`in`)
    val n = scanner.nextInt()
    val cards = mutableListOf<Card>()
    for (i in 0 until n) {
        cards.add(Card(
            scanner.next(),
            scanner.nextInt())
        )
    }

    val originCards = mutableListOf<Card>()
    originCards.addAll(cards)
    quickSort(cards, 0, n - 1)

    println(if (isStable(originCards, cards)) "Stable" else "Not stable")
    for (card in cards) {
        println("${card.pattern} ${card.number}")
    }
}