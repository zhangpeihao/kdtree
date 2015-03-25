kd-tree
======

This kd-tree library just implements create tree and search function,  without add node or remove node functions.

In my project, the kd-tree will become very imbalance when add or remove nodes dynamically.
For the better performance, I need to rebalance the tree frequently.

So, I create this static kd-tree implementation, and it becomes lock-free without modification functions.  The performance is also improved.

I create the new tree by timer in backend thread, the cost is acceptable.
