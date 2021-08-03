import java.util.*

fun countingSort(arrA: IntArray, k: Int): IntArray {
    val n = arrA.size
    val arrC = IntArray(k)
    for (i in 0 until k) arrC[i] = 0
    for (i in 0 until n) arrC[arrA[i]]++
    for (i in 0 until k) arrC[i] = arrC[i] + if (i == 0) 0 else arrC[i - 1]

    val arrB = IntArray(n + 1)
    for (i in 0 until n + 1) arrB[i] = 0
    for (i in n - 1 downTo 0) {
        arrB[arrC[arrA[i]]] = arrA[i]
        arrC[arrA[i]]--
    }
    return arrB
}

fun IntArray.getMax(): Int {
    var max = 0
    for(i in this.indices)
        if (max < this[i]) max = this[i]
    return max
}

fun main(args: Array<String>) {
    val scanner = Scanner(System.`in`)
    val n = scanner.nextInt()
    val arrA = IntArray(n)
    for( i in 0 until n) arrA[i] = scanner.nextInt()

    val arrB = countingSort(arrA, arrA.getMax() + 1)

    println(arrB.joinToString(" ").substring(2))
}