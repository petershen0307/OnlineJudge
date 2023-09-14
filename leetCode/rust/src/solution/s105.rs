pub struct Solution;
use crate::tree_node::node::TreeNode;

// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        // the idea is using recursive find which nodes are belong left subtree and right subtree
        // the first node in preordder is the root of the subtree, we name it as R
        // then use R as the pivot node in inorder, it separate left nodes and right nodes
        if preorder.len() == 0 || inorder.len() == 0 {
            return Option::None;
        }
        // index 0~2999 => -3000 ~ -1, 3000~6000 => 0 ~ 3000
        // value 0: not exist, 1: left, 2: right
        let mut value_map: [i32; 6001] = [0; 6001];
        let mut node = TreeNode::new(preorder[0]);
        let mut left_right = 1;
        let mut left_inorder = Vec::new();
        let mut right_inorder = Vec::new();
        for v in inorder {
            if v == node.val {
                left_right = 2;
            } else {
                value_map[(v + 3000) as usize] = left_right;
                if left_right == 1 {
                    left_inorder.push(v);
                } else {
                    right_inorder.push(v);
                }
            }
        }

        let mut left_preorder = Vec::new();
        let mut right_preorder = Vec::new();
        for v in preorder {
            match value_map[(v + 3000) as usize] {
                1 => {
                    left_preorder.push(v);
                }
                2 => {
                    right_preorder.push(v);
                }
                _ => {}
            }
        }
        node.left = Solution::build_tree(left_preorder, left_inorder);
        node.right = Solution::build_tree(right_preorder, right_inorder);
        return Some(Rc::new(RefCell::new(node)));
    }
}
