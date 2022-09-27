
## Breadth-First Traversal

Given binary tree

```
   0  
  / \ 
 1   2
/ \   
3 4   
````

### Level Order Traversal

How it work:

```
1. Initialize an empty queue.
2. Set root as temp.
3. Loop.
  a. Print data from temp.
  b. Enqueue temp's children (from left to right) to the queue.
  c. Dequeue a node from the queue and set as the temp
```

Result:

```
0 1 2 3 4
```