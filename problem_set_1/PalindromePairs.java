import java.util.*; 

public class PalindromePairs {
    static boolean isPalindrome(String s) {
        StringBuilder reverse = new StringBuilder(s); 
        reverse.reverse(); 

        if(s.equals(reverse.toString())) {
            return true; 
        }

        return false; 
    }

    static ArrayList<int[]> palindromePairs(String[] words) {
        ArrayList<int[]> pairs = new ArrayList<>(); 

        for(int i = 0; i < words.length; i++) {
            for(int j = 0; j < words.length; j++) {
                if(i != j) {
                    String concatinated = words[i] + words[j]; 

                    if(isPalindrome(concatinated)) {
                        int[] pair = {i, j}; 
                        pairs.add(pair); 
                    }
                }
            }
        }

        return pairs; 
    }

    public static void main(String[] args) {
        String[] words = {"tab", "bat", "cat"}; 

        ArrayList<int[]> list = palindromePairs(words); 

        for(int[] pair: list) {
            System.out.print(Arrays.toString(pair) + ", ");
        }
    }
}