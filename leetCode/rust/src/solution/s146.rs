use std::collections::HashMap;

/**
 * Your LRUCache object will be instantiated and called as such:
 * let obj = LRUCache::new(capacity);
 * let ret_1: i32 = obj.get(key);
 * obj.put(key, value);
 */
struct LRUCache {
    map: HashMap<i32, (i32, i32)>, // .0:value, .1:weight
    capacity: usize,
    weight_counter: i32,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl LRUCache {
    fn new(capacity: i32) -> Self {
        LRUCache {
            capacity: capacity as usize,
            map: HashMap::new(),
            weight_counter: 0,
        }
    }

    fn get(&mut self, key: i32) -> i32 {
        match self.map.get_mut(&key) {
            Some(v) => {
                v.1 = self.weight_counter;
                self.weight_counter += 1;
                v.0
            }
            None => -1,
        }
    }

    fn put(&mut self, key: i32, value: i32) {
        // if key existed, just update the value of key
        if let Some(v) = self.map.get_mut(&key) {
            v.0 = value;
            v.1 = self.weight_counter
        } else {
            // key not existed, check capacity first
            if self.map.len() == self.capacity {
                // pop the minimum weight
                let mut minimum_weight = i32::MAX;
                let mut minimum_weight_key: i32 = 0;
                for ele in self.map.iter() {
                    if (ele.1).1 < minimum_weight {
                        minimum_weight = (ele.1).1;
                        minimum_weight_key = *ele.0;
                    }
                }
                self.map.remove(&minimum_weight_key);
            }
            self.map.insert(key, (value, self.weight_counter));
        }
        self.weight_counter += 1;
    }
}

#[test]
fn test_1() {
    let mut cache = LRUCache::new(2);
    cache.put(1, 1);
    cache.put(2, 2);
    assert_eq!(1, cache.get(1));
    cache.put(3, 3);
    assert_eq!(-1, cache.get(2));
    cache.put(4, 4);
    assert_eq!(-1, cache.get(1));
    assert_eq!(3, cache.get(3));
    assert_eq!(4, cache.get(4));
}

#[test]
fn test_2() {
    let mut cache = LRUCache::new(2);
    assert_eq!(-1, cache.get(2));
    cache.put(2, 6);
    assert_eq!(-1, cache.get(1));
    cache.put(1, 5);
    cache.put(1, 2);
    assert_eq!(2, cache.get(1));
    assert_eq!(6, cache.get(2));
}
