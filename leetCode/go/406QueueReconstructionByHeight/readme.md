# Problem 406 Queue Reconstruction By Height
#### Q: Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k),where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.

#### Note:<br/> The number of people is less than 1,100.

## Solution:
* Initial: Stable sort
    1. sorting order: Less than by **k**
        * Example: [[5 0] [7 0] [6 1] [5 2] [7 1] [4 4]]
    1. sorting order: Great than by **h**
        * Example: [[7 0] [7 1] [6 1] [5 0] [5 2] [4 4]]

1. Get the biggest **h** value in input array
1. Insert biggest **h** value to **result** array with **k** as result index
1. repeat

## Note: Every pick has the **biggest h** and the **lowest k**, and the k is the index value in result array
