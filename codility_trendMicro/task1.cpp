#include <iostream>
#include <vector>
using namespace std;

vector<int> to_element(int a)
{
    vector<int> r;
    if (a == 0)
    {
        r.push_back(a);
    }
    else
    {
        for (; a > 0; a /= 10)
        {
            int k = a % 10;
            r.push_back(k);
        }
    }
    return r;
}

int solution1(int A, int B) {
    // write your code in C++11 (g++ 4.8.2)
    if (A > 100000000 || B > 100000000)
    {
        return -1;
    }
    const vector<int>& A1 = to_element(A);
    const vector<int>& B1 = to_element(B);
    int r = 0;
    auto aIter = A1.rbegin();
    auto bIter = B1.rbegin();
    while (true)
    {
        if (aIter != A1.rend())
        {
            r += *aIter;
            r *= 10;
            ++aIter;
        }
        if (bIter != B1.rend())
        {
            r += *bIter;
            r *= 10;
            ++bIter;
        }
        if (aIter == A1.rend() && bIter == B1.rend())
        {
            break;
        }
    }
    r /= 10;
    return r;
}

void print(const vector<int>& a)
{
    for (auto iter = a.rbegin(); iter != a.rend(); ++iter)
    {
        cout << *iter;
    }
}


