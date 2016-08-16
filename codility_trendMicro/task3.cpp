#include <vector>
#include <cassert>
using namespace std;

#define ONE_DAY_PRICE 2
#define SEVEN_DAY_PRICE 7
#define THIRTY_DAY_PRICE 25

int less_than_7days(vector<int> const &days)
{
    int less_7days = 0;
    for (auto ele:days)
    {
        if (ele <= days.front() + 6)
        {
            ++less_7days;
        }
    }
    return less_7days;
}

int solution3(vector<int> &A)
{
    int cost = 0;
    assert(A.size() < 30);
    for (auto ele:A)
    {
        assert(ele <= 30 && ele >= 1);
    }

    for (auto cIter = A.cbegin(); cIter < A.cend();)
    {
        auto tail = cIter + 6 >= A.cend() ? A.cend(): cIter + 6;

        int const less_7days = less_than_7days(vector<int>(cIter, tail));
        if (less_7days >= 4)
        {
            cost += SEVEN_DAY_PRICE;
            cIter += less_7days;
        }
        else
        {
            cost += ONE_DAY_PRICE;
            ++cIter;
        }
    }
    if (cost > THIRTY_DAY_PRICE)
    {
        cost = THIRTY_DAY_PRICE;
    }
    return cost;
}
