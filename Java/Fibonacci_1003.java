import java.io.BufferedReader;
import java.io.InputStreamReader;

public class Fibonacci_1003 {
    private static int T, N;

    public static void main(String[] args) throws Exception {
        int[] F0 = new int[41]; // F0[4] = Fibo(4) 가 호출하는 Fibo(0) 횟수
        int[] F1 = new int[41]; // F1[4] = Fibo(4) 가 호출하는 Fibo(1) 횟수

        F0[0] = F0[2] = F1[1] = F1[2] = 1;
        for(int i=3; i<=40; i++) {
            F0[i] = F0[i-1] + F0[i-2];            
            F1[i] = F1[i-1] + F1[i-2];
        }

        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        T = Integer.parseInt(br.readLine());

        for(int i=0; i<T; i++) {
            N = Integer.parseInt(br.readLine());
            System.out.println(F0[N] + " " + F1[N]);
        }
    }
}
