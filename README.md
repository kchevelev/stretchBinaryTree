# Stretch Binary Tree

stretchBinaryTree is an abstract coding challenge.

## Challenge

Given a binary tree of integer values, stretch the tree vertically by the factor of K, by replacing N nodes with K nodes, with values of K nodes `N / K`, using integer division.

## Conditions
The nodes direction has to be preserved. Left nodes should be streched in the left direction, while right nodes should stretch to the right. Root node, as the node without a parent should stretch to the left.

## Constraints

Do not modify the signature of the stretch interface.

Signature: `stretch(root *TreeNode, stretchAmount int)`

## Original sample tree

```
     12
   /    \
81       34
  \     /  \
   56 19    6
```
## Streched sample tree

```
         6
        /
       6
     /   \
   40     17
  /         \
40          17
 \         /  \
  28      9    3
   \     /      \
    28  9        3

```