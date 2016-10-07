using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace LeetCode
{
    // Definition for a binary tree node.
    class TreeNode
    {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x)
        {
            val = x;
        }
    }

    class SumofLeftLeaves
    {

        public static void Main()
        {
            TreeNode t0 = new TreeNode(0);
            t0.left = new TreeNode(0);
            t0.left.left = new TreeNode(1);
            t0.left.right = new TreeNode(0);
            t0.left.right.left = new TreeNode(1);
            t0.left.right.right = new TreeNode(2);
            t0.right = new TreeNode(2);
            SumofLeftLeaves s = new SumofLeftLeaves();
            Console.WriteLine(s.Solution(t0));
            Console.WriteLine(s.result);
            Console.ReadKey();
        }

        bool IsLeftNodeLeafNode(TreeNode root)
        {
            return null != root &&
                null == root.left && null == root.right;
        }

        int result;

        public int Solution(TreeNode root)
        {
            if (null == root)
            {
                return result;
            }
            if (IsLeftNodeLeafNode(root.left))
            {
                result += root.left.val;
            }
            Solution(root.left);
            Solution(root.right);
            return result;
        }
    }
}
