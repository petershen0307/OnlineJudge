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
        Self::bfs_recursive(root)
    }

    fn bfs_recursive(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
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

    // fn dfs_iter(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
    //     let mut result: Vec<i32> = Vec::new();
    //     let mut current = match root {
    //         Some(n) => n,
    //         None => return result,
    //     };
    //     let mut stack = vec![Rc::clone(&current)];

    //     while !stack.is_empty() {
    //         let tmp = match &(*current).borrow().left {
    //             Some(n) => {
    //                 stack.push(Rc::clone(n));
    //                 // go to left
    //                 Rc::clone(n)
    //             }
    //             None => {
    //                 // in-order put current value to result because current left is empty
    //                 result.push((*current).borrow().val);
    //                 stack.pop(); // because current already in the stack, pop it out
    //                 match &(*current).borrow().right {
    //                     Some(n) => Rc::clone(n),
    //                     None => {
    //                         if !stack.is_empty() {
    //                             stack.pop().unwrap()
    //                         } else {
    //                             break;
    //                         }
    //                     }
    //                 }
    //             }
    //         };
    //         current = tmp;
    //     }
    //     result
    // }
}
