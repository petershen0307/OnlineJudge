int addDigits(int num) {
    int const r = num % 9;
    if (r == 0 && 0 != num)
    {
        return 9;
    }
    return r;
}
