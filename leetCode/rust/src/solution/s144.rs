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
    pub fn preorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        Self::dfs_recursive(root)
    }

    pub fn dfs_recursive(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut result = Vec::new();
        match root.clone() {
            Some(node) => {
                // print to result
                result.push(Rc::clone(&node).borrow().val);
                // go to left
                result.append(&mut Self::dfs_recursive(
                    Rc::clone(&node).borrow().left.clone(),
                ));
                // go to right
                result.append(&mut Self::dfs_recursive(
                    Rc::clone(&node).borrow().right.clone(),
                ));
            }
            None => {}
        }
        result
    }

    fn dfs_iter(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut result = Vec::new();
        let mut stack = Vec::new();
        let mut current = root.clone();
        while !stack.is_empty() || current.is_some() {
            // print current val to result and go to left
            current = match current.clone() {
                Some(node) => {
                    result.push(Rc::clone(&node).borrow().val);
                    stack.push(current.clone());
                    Rc::clone(&node).borrow().left.clone()
                }
                // if current is None, pop node from stack(get parent node) then go to right
                None => {
                    let node = stack.pop().unwrap();
                    // the node pop from stack will not be None, because we didn't push None to stack
                    Rc::clone(&node.unwrap()).borrow().right.clone()
                }
            }
        }
        result
    }
}
