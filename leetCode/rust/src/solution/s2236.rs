// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
pub struct Solution;
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    fn get_real_node(node: Option<Rc<RefCell<TreeNode>>>) -> Option<TreeNode> {
        match node {
            Some(node) => match Rc::try_unwrap(node) {
                Ok(unwrap_node) => Some(unwrap_node.into_inner()),
                Err(_) => None,
            },
            None => None,
        }
    }

    pub fn check_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        match Self::get_real_node(root) {
            Some(node) => {
                let root_value = node.val;
                let l_value: i32 = match Self::get_real_node(node.left) {
                    Some(l_node) => l_node.val,
                    None => 0,
                };
                let r_value: i32 = match Self::get_real_node(node.right) {
                    Some(r_node) => r_node.val,
                    None => 0,
                };
                l_value + r_value == root_value
            }
            None => false,
        }
    }
}
