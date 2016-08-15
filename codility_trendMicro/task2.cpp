#include <iostream>
#include <vector>
#include <string>
#include <cstdio>
#include <utility>
using namespace std;

void print2(const vector<pair<int, int> >& r)
{
    for (auto iter = r.begin(); iter != r.end(); ++iter)
    {
        cout << iter->first << ":" << iter->second << endl;
    }
    cout << "---------" << endl;
}

bool check_time(int t, int c)
{
    if (c == 0)
    {
        if (t > 23)
        {
            return false;
        }
    }
    else
    {
        if (t > 60)
        {
            return false;
        }
    }
    return true;
}

int compare(int a, int b, int c)
{
    int ab = a * 10 + b;
    int bb = b * 10 + a;
    if (ab > bb)
    {
        if (check_time(ab, c))
        {
            return ab;
        }
        else
        {
            return bb;
        }
    }
    if (check_time(bb, c))
    {
        return bb;
    }
    else
    {
        return ab;
    }
    return (a > b) ? a * 10 + b : b * 10 + a;
}

bool check_time(const pair<int, int>& t)
{
    if (t.first > 23 || t.second > 60)
    {
        return false;
    }
    return true;
}

string solution2(int A, int B, int C, int D) {
    // write your code in C++11 (g++ 4.8.2)
    vector<pair<int, int> >r;
    //AB CD
    r.push_back(make_pair(compare(A, B, 0), compare(C, D, 1)));
    //AC BD
    r.push_back(make_pair(compare(A, C, 0), compare(B, D, 1)));
    //AD BC
    r.push_back(make_pair(compare(A, D, 0), compare(B, C, 1)));
    //BC AD
    r.push_back(make_pair(compare(B, C, 0), compare(A, D, 1)));
    //BD AC
    r.push_back(make_pair(compare(B, D, 0), compare(A, C, 1)));
    //CD AB
    r.push_back(make_pair(compare(C, D, 0), compare(A, B, 1)));
//    print2(r);
    for (auto iter = r.begin(); iter != r.end();)
    {
        if (!check_time(*iter))
        {
            iter = r.erase(iter);
        }
        else
        {
            ++iter;
        }
    }
//    print2(r);
    pair<int, int> rr = make_pair(-1, -1);
    for (auto iter = r.cbegin(); iter != r.cend(); ++iter)
    {
        if (iter->first > rr.first)
        {
            rr = *iter;
        }
    }
    if (rr.first == -1)
    {
        return "NOT POSSIBLE";
    }
    char strr[6];
    sprintf(strr, "%02d:%02d", rr.first, rr.second);
    return strr;
}
