# 338. Counting Bits

Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

Example:
For num = 5 you should return [0,1,1,2,1,2].

Follow up:

It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
Space complexity should be O(n).
Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.
Hint:

You should make use of what you have produced already.
Divide the numbers in ranges like [2-3], [4-7], [8-15] and so on. And try to generate new range from previous.

* [problem explain](http://www.cnblogs.com/grandyang/p/5294255.html)
* [C# String.Join()](https://msdn.microsoft.com/zh-tw/library/dd988350(v=vs.110).aspx)
* [C# int array to String array](http://stackoverflow.com/questions/14051257/conversion-from-int-array-to-string-array)