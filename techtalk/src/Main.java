public class Main extends Thread {
    int data = 0;
    String s = "";

    public void run() {
        data++;
    }

    public void compare() {
        if (this.data == 0) {
            this.s = "data:" + this.data;
        }
    }

    public void combine() {
        this.start();
        this.compare();
    }
}

// if (this.s.equals("data:0")) {
// System.out.println("data:0");
// }
// if (this.s.equals("data:1")) {
// System.out.println("data:1");
// }
// if (this.s.equals("")) {
// System.out.println("");
// }