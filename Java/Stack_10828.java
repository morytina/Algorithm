import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.Stack;
import java.util.StringTokenizer;

public class Stack_10828 {
    private static int N;
    public static void main(String[] args) throws Exception {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        N = Integer.parseInt(br.readLine());
        StringTokenizer st = null;

        String cmd; int value;
        Stack<Integer> s = new Stack<Integer>();
        for(int i=0; i<N; i++) {
            st = new StringTokenizer(br.readLine()," ");
            cmd = st.nextToken();
            
            switch (cmd) {
                case "push":
                    value = Integer.parseInt(st.nextToken());
                    s.push(value);
                    break;
                case "pop":
                    if(s.size() == 0)   System.out.println("-1");
                    else                System.out.println(s.pop());
                    break;
                case "size":
                    System.out.println(s.size());
                    break;
                case "empty":
                    if(s.size() == 0)   System.out.println("1");
                    else                System.out.println("0");
                    break;
                case "top":
                    if(s.size() == 0)   System.out.println("-1");
                    else                System.out.println(s.peek());
                    break;
                default:
                    break;
            }
        }
    }
}
