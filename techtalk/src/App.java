public class App {
    public static void main(String[] args) throws Exception {
        for (int i = 0; i < 1_000_000; i++) {
            Main thread = new Main();
            thread.combine();
        }
    }
}
