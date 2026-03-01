package main

class Solution {
    fun isAnagram(
        s: String,
        t: String,
    ): Boolean {
        if (s.length != t.length) return false

        val map = mutableMapOf<Char, Int>()

        for (char in s) {
            map[char] = map.getOrDefault(char, 0) + 1
        }

        for (char in t) {
            val count = map.getOrDefault(char, 0) - 1
            if (count < 0) return false
            map[char] = count
        }

        return true
    }
}
