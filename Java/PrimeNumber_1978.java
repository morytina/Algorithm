import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class PrimeNumber_1978 {
    private static int N;
    private static boolean[] isPrime;

    public static void main(String[] args) throws Exception {
        makePrime();
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        N = Integer.parseInt(br.readLine());

        StringTokenizer st = new StringTokenizer(br.readLine()," ");
        
        int ans = 0;
        for(int i=0; i<N; i++) {
            int num = Integer.parseInt(st.nextToken());
            if(isPrime[num]) ans++;
        }
        System.out.println(ans);
    }

    public static void makePrime() {
        isPrime = new boolean[1001];
        for(int i=2; i<1001; i++) isPrime[i] = true;
        for(int i=2; i<1001; i++) {
            if(isPrime[i]) {
                int cur = i << 1;
                while(cur < 1001) {
                    isPrime[cur] = false;
                    cur += i;
                }
            }
        }
    }
}
