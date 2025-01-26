import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Average_1546 {
    private static int N;

    public static void main(String[] args) throws Exception {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        N = Integer.parseInt(br.readLine());
        StringTokenizer st = new StringTokenizer(br.readLine()," ");

        float[] score = new float[N];
        float max = 0;
        for(int i=0; i<N; i++) {
            score[i] = Float.parseFloat(st.nextToken());
            if(score[i] > max) max = score[i];
        }

        float sum = 0;
        for(int i=0; i<N; i++) {
            score[i] = score[i] / max * 100;
            sum += score[i];
        }

        System.out.println(sum / N);
    }
}
