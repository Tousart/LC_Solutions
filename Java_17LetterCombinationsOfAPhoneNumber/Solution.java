package LC17_LetterCombinationsOfAPhoneNumber;

import java.util.ArrayList;
import java.util.List;

class Solution {
    public static String[] lettersCombinations = {
            "abc", "def",
            "ghi", "jkl",
            "mno", "pqrs",
            "tuv", "wxyz"};

    public List<String> letterCombinations(String digits) {
        List<String> result = new ArrayList<>();

        if (digits.isEmpty()) {
            return result;
        }

        func(0, new StringBuilder(), digits, result);

        return result;
    }

    public static void func(int index, StringBuilder elem, String digits, List<String> result) {
        if (index == digits.length()) {
            result.add(elem.toString());
            return;
        }

        String letters = lettersCombinations[digits.charAt(index) - '0' - 2];

        for (char letter : letters.toCharArray()) {
            func(index + 1, elem.append(letter), digits, result);
            elem.deleteCharAt(elem.length() - 1);
        }
    }
}
