package LC_DivideTwoIntegers;

public class Solution {
    public int divide(int dividend, int divisor) {

        if (dividend == Integer.MIN_VALUE && divisor == -1)
            return Integer.MAX_VALUE;

        long absDividend = Math.abs((long)dividend);
        long absDivisor = Math.abs((long)divisor);
        int result = 0;

        while (absDividend >= absDivisor) {
            int shift = 0;

            while (absDividend >= (absDivisor << shift)){
                shift++;
            }

            System.out.println(shift);
            result += 1 << (shift - 1);
            absDividend -= (absDivisor << (shift - 1));
            System.out.println(result);
            System.out.println(absDividend + "\n");
        }

        if (dividend > 0 && divisor < 0 ||
                dividend < 0 && divisor > 0)
            return -result;

        return result;
    }
}