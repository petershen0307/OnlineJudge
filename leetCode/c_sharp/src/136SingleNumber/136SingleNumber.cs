using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode
{
    class _136SingleNumber
    {
        static void Main()
        {
            int[] nums = {1, 2, 3, 2, 1};
            _136SingleNumber r = new _136SingleNumber();
            Console.WriteLine(r.SingleNumber(nums));
            Console.ReadKey();
        }

        // A xor A == 0
        public int SingleNumber(int[] nums)
        {
            int result = nums[0];
            // lamda function: (num, index) => 0 != index
            // argument: num, index
            // return: 0 != index
            foreach (int i in nums.Where((num, index) => 0 != index))
            {
                result ^= i;
            }
            return result;
        }
    }
}
