pub struct Solution;
use crate::tree_node::node::TreeNode;
use std::borrow::BorrowMut;
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
    pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        Self::dfs_recursive(root)
    }

    pub fn dfs_recursive(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut result = Vec::new();
        match root.clone() {
            Some(node) => {
                result.append(&mut Self::dfs_recursive(
                    Rc::clone(&node).borrow().left.clone(),
                ));
                result.append(&mut Self::dfs_recursive(
                    Rc::clone(&node).borrow().right.clone(),
                ));
                result.push(Rc::clone(&node).borrow().val);
            }
            None => {}
        }
        result
    }

    fn dfs_iter(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        // https://www.enjoyalgorithms.com/blog/iterative-binary-tree-traversals-using-stack
        let mut result = Vec::new();
        let mut stack = Vec::new();
        let mut right_child_stack = Vec::new();
        let mut current = root.clone();
        while !stack.is_empty() || current.is_some() {
            current = match current.clone() {
                Some(node) => {
                    if Rc::clone(&node).borrow().right.is_some() {
                        right_child_stack.push(Rc::clone(&node).borrow().right.clone());
                    }
                    stack.push(current.clone());
                    Rc::clone(&node).borrow().left.clone()
                }
                None => {
                    let node = stack.last().unwrap();
                    if !right_child_stack.is_empty()
                        && Rc::clone(node.as_ref().unwrap()).borrow().right
                            == *right_child_stack.last().unwrap()
                    {
                        right_child_stack.pop().unwrap()
                    } else {
                        result.push(Rc::clone(node.as_ref().unwrap()).borrow().val);
                        stack.pop().unwrap();
                        None
                    }
                }
            }
        }
        result
    }
}
