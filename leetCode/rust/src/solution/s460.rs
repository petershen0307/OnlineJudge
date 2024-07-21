use std::collections::HashMap;

struct CacheContent {
    value: i32,
    frequency: i32,
    order: i32,
}
struct LFUCache {
    capacity: usize,
    mono_counter: i32,
    map: HashMap<i32, CacheContent>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl LFUCache {
    fn new(capacity: i32) -> Self {
        LFUCache {
            capacity: capacity as usize,
            mono_counter: 1,
            map: HashMap::new(),
        }
    }

    fn get(&mut self, key: i32) -> i32 {
        if let Some(v) = self.map.get_mut(&key) {
            v.frequency += 1;
            v.order = self.mono_counter;
            self.mono_counter += 1;
            v.value
        } else {
            -1
        }
    }

    fn put(&mut self, key: i32, value: i32) {
        // if key existed update the frequency
        if let Some(v) = self.map.get_mut(&key) {
            v.frequency += 1;
            v.value = value;
            return;
        }
        if self.map.len() == self.capacity {
            // remove least used item
            let mut f = std::i32::MAX;
            let mut order = std::i32::MAX;
            let mut item_key: Option<i32> = None;
            for (key, v) in self.map.iter() {
                if v.frequency < f {
                    f = v.frequency.to_owned();
                    order = v.order;
                    item_key = Some(key.to_owned());
                }
                if v.frequency == f && v.order < order {
                    f = v.frequency.to_owned();
                    order = v.order;
                    item_key = Some(key.to_owned());
                }
            }
            self.map.remove(&item_key.unwrap());
        }
        self.map.insert(
            key,
            CacheContent {
                value,
                frequency: 1,
                order: self.mono_counter,
            },
        );
        self.mono_counter += 1;
    }
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * let obj = LFUCache::new(capacity);
 * let ret_1: i32 = obj.get(key);
 * obj.put(key, value);
 */

#[test]
fn test_1() {
    let mut lfu = LFUCache::new(2);
    lfu.put(1, 1); // cache=[1,_], cnt(1)=1
    lfu.put(2, 2); // cache=[2,1], cnt(2)=1, cnt(1)=1
    assert_eq!(1, lfu.get(1)); // return 1
                               // cache=[1,2], cnt(2)=1, cnt(1)=2
    lfu.put(3, 3); // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
                   // cache=[3,1], cnt(3)=1, cnt(1)=2
    assert_eq!(-1, lfu.get(2)); // return -1 (not found)
    assert_eq!(3, lfu.get(3)); // return 3
                               // cache=[3,1], cnt(3)=2, cnt(1)=2
    lfu.put(4, 4); // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
                   // cache=[4,3], cnt(4)=1, cnt(3)=2
    assert_eq!(-1, lfu.get(1)); // return -1 (not found)
    assert_eq!(3, lfu.get(3)); // return 3
                               // cache=[3,4], cnt(4)=1, cnt(3)=3
    assert_eq!(4, lfu.get(4)); // return 4
                               // cache=[4,3], cnt(4)=2, cnt(3)=3
}

#[test]
fn test_2() {
    let mut lfu = LFUCache::new(2);
    lfu.put(2, 1);
    lfu.put(3, 2);
    assert_eq!(2, lfu.get(3));
    assert_eq!(1, lfu.get(2));

    lfu.put(4, 3);
    assert_eq!(1, lfu.get(2));
    assert_eq!(-1, lfu.get(3));
    assert_eq!(3, lfu.get(4));
}
