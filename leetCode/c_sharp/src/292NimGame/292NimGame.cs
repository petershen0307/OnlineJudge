using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode
{
    class _292NimGame
    {
        static void Main()
        {
            _292NimGame t = new _292NimGame();
            Console.WriteLine(t.CanWinNim(100));
            Console.ReadKey();
        }

        public bool CanWinNim(int n)
        {
            return n % 4 != 0;
        }
    }
}
