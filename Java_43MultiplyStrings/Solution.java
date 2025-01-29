package LC43_MultiplyStrings;

public class Solution {
    public String multiply(String num1, String num2) {
        if (num1.equals("0") || num2.equals("0"))
            return "0";

        StringBuilder result = new StringBuilder();
        int[] numbers = new int[num1.length() + num2.length()];

        for (int i = num1.length() - 1; i >= 0; i--) {
            for (int j = num2.length() - 1; j >= 0; j--) {
                int num = (num1.charAt(i) - '0') * (num2.charAt(j) - '0');
                numbers[i + j + 1] += num;
            }
        }

        for (int i = numbers.length - 1; i >= 0; i--) {

            if (i == 0 && numbers[i] == 0) {
                break;
            }

            if (numbers[i] / 10 != 0) {
                result.insert(0, numbers[i] % 10);
                numbers[i - 1] += numbers[i] / 10;
            } else {
                result.insert(0, numbers[i]);
            }
        }

        return result.toString();
    }
}
