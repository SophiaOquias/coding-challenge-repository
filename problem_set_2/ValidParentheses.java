import java.util.*; 

public class ValidParentheses {
    static boolean isOpenBracket(char bracket) {
        if(bracket == '(' || bracket == '[' || bracket == '{') {
            return true; 
        }

        return false; 
    }

    static char getMatch(char bracket) {
        switch (bracket) {
            case ')': return '(';
            case ']': return '['; 
            case '}': return '}';
            default: return 0;
        }
    }

    static boolean isValid(String s) {
        Stack<Character> stack = new Stack<>(); 

        for(int i = 0; i < s.length(); i++) {
            char current = s.charAt(i); 

            if(isOpenBracket(current)) {
                stack.push(current); 
            }
            else {
                if(stack.isEmpty()) {
                    return false; 
                }

                if(stack.peek() != getMatch(current)) {
                    return false; 
                }

                stack.pop(); 
            }
        }

        return true; 
    }

    public static void main(String[] args) {
        System.out.println(isValid("([)]")); // false
    }
}