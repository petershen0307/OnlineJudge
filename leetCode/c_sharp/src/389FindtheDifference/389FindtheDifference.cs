using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode
{
    class _389FindtheDifference
    {
        static void Main()
        {
            _389FindtheDifference t = new _389FindtheDifference();
            Console.WriteLine(t.FindTheDifference2("avb", "bbva"));
            Console.ReadKey();
        }

        public char FindTheDifference(string s, string t)
        {
            string totalStr = s + t;
            char result = totalStr[0];
            foreach (char i in totalStr.Where((num, index) => 0 != index))
            {
                result = (char)(result ^ i);
            }
            return result;
        }

        public char FindTheDifference2(string s, string t)
        {
            Dictionary<char, int> charCount = new Dictionary<char, int>();
            string totalStr = s + t;
            foreach (char i in totalStr)
            {
                if (!charCount.ContainsKey(i))
                {
                    charCount[i] = 0;
                }
                charCount[i]++;
            }
            char result = '\0';
            foreach (KeyValuePair<char, int> i in charCount)
            {
                if (i.Value % 2 == 1)
                {
                    result = i.Key;
                    break;
                }
            }
            return result;
        }
    }
}
