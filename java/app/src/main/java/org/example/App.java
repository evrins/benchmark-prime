/*
 * This source file was generated by the Gradle 'init' task
 */
package org.example;


public class App {

    public static boolean isPrime(long n) {
        if (n < 2) {
            return false;
        }

        if (n == 2) {
            return true;
        }

        long max = ((long) Math.ceil(Math.sqrt(n)));

        for (int i = 2; i <= max; i++) {
            if (n % i == 0) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) {
        if (args.length != 1) {
            System.out.printf("Usage: java -jar App.jar <number>%n");
            System.exit(1);
        }
        long n = Long.parseLong(args[0]);
        long start = System.currentTimeMillis();
        long count = 0;

        for (int i = 0; i < n; i++) {
            if (isPrime(i)) {
                count++;
            }
        }

        long end = System.currentTimeMillis();
        System.out.printf("java, %d, %d, %d%n", n, count, end - start);
    }
}
