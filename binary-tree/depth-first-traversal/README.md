## Depth-First Traversal

Given binary tree

```
   0  
  / \ 
 1   2
/ \   
3 4   
````

### Pre Order Traversal

How it work:

```
1. Traverse root.
2. Pre-order on left sub-tree.
3. Pre-order on right sub-tree.
```

Result:

```
0 1 3 4 2
```

### In Order Traversal

How it work:

```
1. In-order on left sub-tree.
2. Traverse root.
3. In-order on right sub-tree.
```

Result:

```
3 1 4 0 2
```

### Post Order Traversal

How it work:

```
1. Post-order on left sub-tree.
2. Post-order on right sub-tree.
3. Traverse root.
```

Result:

```
3 4 1 2 0
```
