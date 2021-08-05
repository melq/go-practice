import java.util.*

fun swap(arr: IntArray, i: Int, j: Int) {
   val tmp = arr[i]
   arr[i] = arr[j]
   arr[j] = tmp
}

fun partition(arr: IntArray, p: Int, r: Int): Int {
   val x = arr[r]
   var i = p - 1
   for (j in p until r) {
      if (arr[j] <= x) {
         i++
         swap(arr, i, j)
      }
   }
   swap(arr, i + 1, r)
   return i + 1
}

fun main(args: Array<String>) {
   val scanner = Scanner(System.`in`)
   val n = scanner.nextInt()
   val arr = IntArray(n)
   for (i in 0 until n) arr[i] = scanner.nextInt()

   val res = partition(arr, 0, n - 1)
   print("${arr.slice(0 until res).joinToString(" ")} [${arr[res]}] ${arr.slice(res + 1 until n).joinToString(" ") }\n")
}