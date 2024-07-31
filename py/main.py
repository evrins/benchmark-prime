import time
import math
import sys


def is_prime(n: int) -> bool:
    if n < 2:
        return False

    if n == 2:
        return True

    max = math.ceil(math.sqrt(n)) + 1

    for i in range(2, max):
        if n % i == 0:
            return False

    return True


def main(n):
    count = 0
    start_time = time.time()
    for i in range(0, n):
        if is_prime(i):
            count += 1

    end_time = time.time()

    print(f'py, {n}, {count}, {round((end_time - start_time) * 1000)}')


if __name__ == '__main__':
    if len(sys.argv) != 2:
        print(f'Usage: {sys.argv[0]} <n>')
        sys.exit(1)
    main(int(sys.argv[1]))
