import collections


def nthSuperUglyNumber(n, primes):
    total_number_count = 1
    current_count = 1
    non_ugly_numbers = set()
    ugly_numbers = set([1])
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
                ugly_numbers.add(current_count)
            else:
                non_ugly_numbers.add(current_count)
    print(sorted(ugly_numbers))
    return current_count
