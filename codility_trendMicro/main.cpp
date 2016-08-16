#include <iostream>
#include <string>
#include <vector>
using namespace std;

int solution1(int A, int B);
string solution2(int A, int B, int C, int D);
int solution3(vector<int> &A);

int main()
{
    // 123456789
    //print(to_element(0));
    //cout << solution1(78, 100000000);
    //cout << solution2(1, 2, 3, 4) << endl;
    //cout << solution2(9, 9, 9, 9) << endl;
    //cout << solution2(0, 2, 0, 0) << endl;
    vector<int> A = {1, 2, 4, 5, 7, 29, 30};
    cout << solution3(A);
    return 0;
}
