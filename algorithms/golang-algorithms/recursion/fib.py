import time


def time_exec(fn):
    start_time = time.time()
    fn()
    end_time = time.time()
    print(f"Elapsed time: {end_time - start_time} seconds")


def fib(n):
    if n < 2:
        return n

    return fib(n - 1) + fib(n - 2)


def fib2(n: int) -> int:
    cache = {}
    return _helper(n, cache)


def _helper(n: int, cache: dict) -> int:
    if n < 2:
        return n

    if n in cache:
        return cache[n]

    res = _helper(n - 1, cache) + _helper(n - 2, cache)
    cache[n] = res

    return res


def fib3(n):
    if n < 2:
        return n

    a = 0
    b = 1

    for _ in range(2, n, 1):
        a, b = b, a + b

    return b


def main():
    fibTarget = 35
    time_exec(lambda: print(fib(fibTarget)))
    time_exec(lambda: print(fib2(fibTarget)))
    time_exec(lambda: print(fib3(fibTarget)))


if __name__ == "__main__":
    main()
