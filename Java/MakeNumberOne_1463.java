import java.io.BufferedReader;
import java.io.InputStreamReader;

public class MakeNumberOne_1463 {
    private static int N;
    private static int[] D = new int[1000001]; // D[a] = 1에서 시작해 a 를 만들 수 있는 최소 연산횟수 (*2, *3, +1) 
    public static void main(String[] args) throws Exception {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        N = Integer.parseInt(br.readLine());

        D[2] = D[3] = 1; // init
        for(int i=4; i<=1000000; i++) {
            D[i] = D[i-1] + 1;
            if(i%2 == 0) D[i] = Math.min(D[i], D[i/2] + 1);
            if(i%3 == 0) D[i] = Math.min(D[i], D[i/3] + 1);
        }

        System.out.println(D[N]);
    }
}
