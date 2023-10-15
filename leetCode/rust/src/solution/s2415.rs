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
    pub fn reverse_odd_levels(
        root: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        use std::collections::VecDeque;
        let mut queue: VecDeque<Rc<RefCell<TreeNode>>> = VecDeque::new();
        if root.is_none() {
            return None;
        }
        let mut level = 0;
        match root.as_ref() {
            Some(n) => {
                queue.push_back(Rc::clone(&n));
            }
            None => {
                return None;
            }
        }

        while !queue.is_empty() {
            // pop all nodes from queue
            for _ in 0..1 << level {
                match queue.pop_front() {
                    None => (),
                    Some(n) => {
                        // then put every node's children to the queue
                        match (*n).borrow().left.clone() {
                            Some(l) => {
                                queue.push_back(Rc::clone(&l));
                            }
                            None => (),
                        }
                        match (*n).borrow().right.clone() {
                            Some(r) => {
                                queue.push_back(Rc::clone(&r));
                            }
                            None => (),
                        }
                    }
                }
            }
            // finish last level node in the queue
            level += 1;
            // check current level is odd then reverse all node value in current queue
            if level % 2 == 1 {
                for i in 0..queue.len() / 2 {
                    let temp = queue[i].borrow().val;
                    let temp2 = queue[queue.len() - i - 1].borrow().val;
                    (*queue[i]).borrow_mut().val = temp2;
                    (*queue[queue.len() - i - 1]).borrow_mut().val = temp;
                }
            }
        }
        root
    }
}

#[test]
fn test_1() {
    let mut root = TreeNode::new(1);
    root.left = Some(Rc::new(RefCell::new(TreeNode::new(2))));
    root.right = Some(Rc::new(RefCell::new(TreeNode::new(3))));
    Solution::reverse_odd_levels(Some(Rc::new(RefCell::new(root))));
}
