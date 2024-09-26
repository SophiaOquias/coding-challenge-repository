import java.util.*; 

public class LongestIncreasingSubsequence {

    static int lengthOfLIS(int[] num) {
        int longestLength = 0; 

        for(int i = 0; i < num.length; i++) {
            int current = num[i]; 
            int sum = 0; 

            for(int j = i; j < num.length; j++) {
                if(current <= num[j]) {
                    sum++; 
                    current = num[j]; 
                }

                if(j == num.length - 1 && longestLength < sum) {
                    longestLength = sum; 
                }
            }
        }

        return longestLength; 
    }

    public static void main(String[] args) {
        int[] inputNumbers = {10, 9, 2, 5, 3, 7, 101, 18};
        System.out.println(lengthOfLIS(inputNumbers));
    }
}