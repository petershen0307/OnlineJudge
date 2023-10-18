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
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        Self::dfs_recursive(root)
    }

    fn dfs_recursive(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut result: Vec<i32> = Vec::new();
        match root {
            None => return result,
            Some(n) => {
                // go to left
                result.append(&mut Solution::inorder_traversal((*n).borrow().left.clone()));
                // add to result
                result.push((*n).borrow().val);
                // go to right
                result.append(&mut Solution::inorder_traversal(
                    (*n).borrow().right.clone(),
                ));
            }
        }
        result
    }

    fn dfs_iter(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        // https://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion/
        let mut result: Vec<i32> = Vec::new();
        let mut current = match root.as_ref() {
            Some(_) => root.clone(),
            None => return result,
        };
        let mut stack = Vec::new();

        while !stack.is_empty() || !current.is_none() {
            current = match current.clone() {
                // push current to stack then go to left
                Some(node) => {
                    stack.push(current.clone());
                    (*node).borrow().left.clone()
                }
                // pop node from stack and print the result then go to right
                None => {
                    let node = stack.pop();
                    let val = node.clone();
                    result.push((*val.unwrap().unwrap()).borrow().val);
                    (*node.unwrap().unwrap()).borrow().right.clone()
                }
            };
        }
        result
    }
}
