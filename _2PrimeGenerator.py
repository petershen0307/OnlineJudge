import math


def is_prime_number(number):
    assert 0 < number
    if 1 == number:
        return False
    root_int = math.floor(math.sqrt(number))
    for i in range(2, root_int + 1):
        if number % i == 0:
            return False
    return True


def enhance_prime_detect(number, prime_numbers):
    if len(prime_numbers) == 0:
        return is_prime_number(number)
    else:
        root_int = math.floor(math.sqrt(number))
        for i in prime_numbers:
            if i > root_int:
                break
            elif number % i == 0:
                return False
    return True


def sieve_of_eratosthenes(head, tail):
    all_range = [i for i in range(2, tail + 1)]
    index = 0
    while index < len(all_range):
        current = all_range[0]
        all_range = [i for i in all_range if i % current != 0]
        all_range.append(current)
        index += 1
    return [i for i in all_range if i >= head]


def prime_generator(head, tail):
    assert 0 < head and tail > 0 and tail >= head
    prime_numbers = []
    for number in range(1, tail + 1):
        if enhance_prime_detect(number, prime_numbers):
            prime_numbers.append(number)
    prime_numbers = [i for i in prime_numbers if i >= head]
    return prime_numbers


def start():
    times = input()
    input_list = []
    for i in range(int(times)):
        input_list.append(list(map(int, input().split())))
    for input_str in input_list:
        head = int(input_str[0])
        tail = int(input_str[1])
        result = prime_generator(head, tail)
        print('\n'.join(map(str, result)))
        print('\n', end='')

import time
if __name__ == '__main__':
    start_time = time.time()
    print('start time {0}'.format(start_time))
    #prime_generator(1, 5000000)
    print(sieve_of_eratosthenes(1, 100000))
    print('execution time {result}'.format(result=(time.time() - start_time)))
