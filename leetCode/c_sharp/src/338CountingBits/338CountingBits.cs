using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace LeetCode
{
    class _338CountingBits
    {
        static void Main()
        {
            _338CountingBits test = new _338CountingBits();
            Console.WriteLine(test.CountBits(5));
        }
        // 指數 ~1 value 0 1
        // 指數1~2 value 2 3
        // 指數2~3 value 4 5 6 7
        // 指數3~4 value 8 9 10 11 12 13 14 15
        public int[] CountBits(int num)
        {
            int[] result = new int[num+1];
            result[0] = 0;
            if (num >= 1)
            {
                result[1] = 1;
            }
            for (int e = 1; 1<<e <= num; e++)
            {
                // (1 << e) == 2^e
                int end = num < ((1 << (e + 1)) - 1) ? num : ((1 << (e + 1)) - 1);
                int start = (1 << e);
                for (int i = start; i <= end; i++)
                {
                    result[i] = result[i - start];
                }
            }
            return result;
        }
    }
}
