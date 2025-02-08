package Kotlin_2349DesignANumberContainerSystem

import java.util.SortedSet

/**
 * Your NumberContainers object will be instantiated and called as such:
 * var obj = NumberContainers()
 * obj.change(index,number)
 * var param_2 = obj.find(number)
 */

class NumberContainers() {
    private val indexDict: MutableMap<Int, Int> = mutableMapOf()
    private val numberDict: MutableMap<Int, SortedSet<Int>> = mutableMapOf()

    fun change(index: Int, number: Int) {
        val currentNumber = indexDict[index]

        if (currentNumber != null) {
            numberDict[currentNumber]?.remove(index)
            // Если такого числа больше нет ни в одном из контейнеров
            if (numberDict[currentNumber]?.isEmpty() == true) {
                numberDict.remove(currentNumber)
            }
        }

        indexDict[index] = number
        numberDict.computeIfAbsent(number) { sortedSetOf() }.add(index)
    }

    fun find(number: Int): Int {
//        println(indexDict)
//        println(numberDict)
        return numberDict[number]?.first() ?: -1
    }

}



fun main() {
    val sl = NumberContainers()
    sl.change(2, 5)
    sl.change(1, 6)
    sl.change(3, 7)
    sl.change(1, 7)
    println(sl.find(6))
}