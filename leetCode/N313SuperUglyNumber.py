import sys
import heapq


def method1(n, primes):
    """
    :type n: int
    :type primes: List[int]
    :rtype: int
    """
    total_number_count = 1
    current_count = 1
    non_ugly_numbers = set()
    while total_number_count != n:
        current_count += 1
        is_non_ugly_number = False
        for non_ugly_num in non_ugly_numbers:
            if current_count % non_ugly_num == 0:  # non ugly number
                is_non_ugly_number = True
                non_ugly_numbers.add(current_count)
                break
        if not is_non_ugly_number:
            is_ugly_number = False
            for prime in primes:
                if current_count % prime == 0:  # ugly number
                    is_ugly_number = True
                    break
            if is_ugly_number:
                total_number_count += 1
            else:
                non_ugly_numbers.add(current_count)
    return current_count


# O(k*n) k is len(primes)
def method2(n, primes):
    """
    :type n: int
    :type primes: List[int]
    :rtype: int
    """
    ugly_list = [1] + [sys.maxsize] * (n - 1)
    k = len(primes)
    index = [0] * k
    for i in range(1, n):
        added_index = []
        for j in range(k):
            if primes[j] * ugly_list[index[j]] <= ugly_list[i]:
                ugly_list[i] = primes[j] * ugly_list[index[j]]
                added_index.append(j)
                # ugly_list[i] = min(ugly_list[i], primes[j] * ugly_list[index[j]])
        for ele in added_index:
            if ugly_list[i] == (primes[ele] * ugly_list[index[ele]]):
                index[ele] += 1
                # for j in range(k):
                #    if ugly_list[i] == (primes[j] * ugly_list[index[j]]):
                #        index[j] += 1
    return ugly_list[-1]


def method3(n, primes):
    uglys = [1]

    def gen_ugly(prime):
        for ele in uglys:
            yield ele * prime

    merge = heapq.merge(*map(gen_ugly, primes))
    while len(uglys) < n:
        ugly = next(merge)
        if uglys[-1] != ugly:
            uglys.append(ugly)
    return uglys[-1]


if __name__ == "__main__":
    def square(n):
        return n ** 2


    print(*map(square, range(1, 10)))
    method3(12, [2, 7, 13, 19])
