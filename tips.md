# 提示

## 3, [无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

双指针表示滑动窗口

用**快慢指针**分别维护一个子串的窗口，并把字串中存在的字符，用一个 map 来保存。如果**快**指针指向的字符并不存在于 map 中，那么保存此字符，快指针向右移动，重新计算最大窗口的大小。否则就将**慢**指向的字符从 map 中删除，并向右移动慢指针。如此循环直到快或者慢指针达到字符串末尾。

最坏情况下，快慢指针分别遍历字符串一次，时间复杂度为 O(n)。map 保存所有字符，空间复杂对为 O(n)

在上面的方法中，慢指针一次只能移动一位，不如在遍历字符串的过程中，用 map 来存储已出现的每个字符的 index。这样一来，一旦遇到一个重复的字符，直接就从 map 中读出这个字符上次出现的位置，并把 slow 移动到这个位置。在遍历的同时update 最大窗口的长度。

M15, [三数之和](https://leetcode-cn.com/problems/3sum/)

先排序，遍历数组 nums 固定第一个数，下标为 i。然后第二个数的下标j = i+1, 第三个数下标 k = len(nums)-1。如果三个数和 > 0, k--; <0, j++; ==0, 记录答案。

- 注意遍历数组的时候，如果和上一个数相等就要跳过。sum == 0 的时候，如果 j 和 k 的下一个元素也相等，也要跳过。否则都可能生成重复答案。

- 注意移动 j 和 k 的时候的边界

## 23, [合并 K 个有序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

分治：

```go
merge2Lists := func(l1, l2 *ListNode) *ListNode 
...
left := lists[:middle]
right := lists[middle:]
return merge2Lists(mergeKLists(left), mergeKLists(right))
```

## 98, [验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)

1，中序遍历

BST 的中序遍历可以得到一个升序的有序数组。检查数组的有序性就可以验证BST。

时间复杂度是 O(n)，因为要遍历所有n个节点。空间复杂度也是 O(n), 要用数组存储所有n个节点的值来做验证。

2, 迭代

主要是要在迭代函数中传入节点值的上限(max)和下限(min)。max 是用来限制左子树节点中的节点最大值，并在迭代过程中，用当前节点的 val 作为 max 值。反之右子树亦然。

时间复杂度 O(n), 空间复杂度 O(n)

## 109, [有序链表转换BST](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/)

写一个函数用快慢指针的方法找到中间节点，注意循环的条件: fast != right && fast.Next != right。把中点作为根，用递归来构建根的左子树和右子树，其中停止条件是链表的 **头节点 == 尾节点**。

## 114, [二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)

分解为递归子问题：

- 如果节点为 nil，直接返回。

- 否则先展开节点的左孩子，再展开右孩子。

- 然后将节点左孩子链表接到节点的右孩子指针上，遍历到链表的末尾，再把右孩子的链表接上。

- 注意接右孩子的同时，把左孩子指针设置为 nil

## 136, [只出现一次的数字I](https://leetcode-cn.com/problems/single-number/)

题目：给定一个**非空**整数数组，除了某个元素只出现**一次**以外，其余每个元素均出现**两次**。找出那个只出现了一次的元素。

哈希：创建map，遍历数组作为 key，如果 key 不存在就加入 map，已存在就 delete。到最后 map 中唯一的 key 就是出现一次的元素。

排序：将数组排序，然后遍历数组，找出与前后都不一样的元素。

集合：将所有元素放入 set 去重，然后计算 sum*2，减去原数组中所有的元素得到答案。

位运算：任何数和 0 做异或运算，结果仍然是原来的数；任何数和其自身做异或运算，结果是 0；异或运算满足交换律和结合律

```go
func singleNumber4(nums []int) (single int) {
    for _, num := range nums {
        single ^= num
    }
    return
}
```

## 146, [实现 LRU 缓存](https://leetcode-cn.com/problems/lru-cache/)

双向链表 + 哈希表

1, 在链表节点结构的定义中，加入 key 字段，方便从哈希表中删除对应节点

2, Get, Put 操作会把节点移到链表尾端。Put 操作会添加节点到尾端，还可能会删除第一个节点。将这些频繁的链表操作实现为单独的方法，可以让 Get 和 Put 的代码更简洁清晰。

## 147, [对链表进行插入排序](https://leetcode-cn.com/problems/insertion-sort-list/)

外层 for 循环遍历未排序节点，内层 for 进行插入排序。时间复杂度为 O(n^2), 空间复杂度为 O(1)。最后注意下面两行代码循序不要搞反：

```go
curr.Next = currSorted.Next
currSorted.Next = curr
```

## 148, [排序链表](https://leetcode-cn.com/problems/sort-list/)

用归并排序。创建 merge 函数合并两个有序链表。创建 cut 函数将一个链表切成两个链表，返回的第一个链表尾部要改为 nil。递归实现 sortList, 对于三个节点以上的链表，先 cut，再递归调用 sortList 排序左右子链表，再合并已排序的链表。对于只有两个节点的链表，用交换两个节点的 Val 来实现排序更方便。

## 239, [滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)

如果数组长度是L，窗口长度是K，暴力查找的时间复杂度为 O(LK)，是不能接受的。优化点在于减少每次在窗口中查找最大数的时间。

方法之一是用双向队列维护一个当前窗口中最大数的列表，最大的数字永远在第 0 个位置。

维护这样一个队列的方法：

遍历整个数组，当前数字为 n，下标为 i：

```text
1，从后向前遍历队列 queue，如果数字小于n，则将其从队列中移除。

2，然后将 n 放入队列末尾。到此为止的目的是不但可以保证最大数字在队列开头，而且将来这个数字被移出窗口后，后面第二大的数字可以马上顶上。

3，检查当前窗口是否已经开始右移 (i>=k)，而且被从左边移出去的数字正好等于队列中的最大数 (nums[i-k] == queue[0]), 如果是，那么就得把它从队列移出，然后队列中第二大的数就会变成最大的数，出现在 0 位置。

4，如果 i 已经和窗口一样大了 (i >= k-1)，那么就得开始填写答案了: res[i-k+1] = queue[0]
```

主要注意几个操作的临界条件：

```text
1，窗口开始右移：i >= k; 此次右移从左边被移出窗口的数字：nums[i-k]

2， 开始填写答案的条件 i >= k-1; 此次答案的下标： i-k+1
```

## 350, [两个数组的交集II](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)

先排序两个数组，用双指针。复用 nums1 数组的空间保存返回值。

## 387, [字符串中的第一个唯一字符](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)

- 第一次遍历字符串，用 map 保存每个字符最后出现的位置。第二次遍历字符串，同时查找 map，下标相同就返回答案, 不同就改为 -1 以免下次匹配上

- 可以用 [26]int 来代替 map

## 496, [下一个更大元素I](https://leetcode-cn.com/problems/next-greater-element-i/)

单调栈：

- 反向遍历 nums2, 如果当前数比栈顶的数大，就移除栈顶元素，直到遇到比它大的数，或者栈空

- 在把当前数入栈前，用 map 保存其作为 key，栈顶元素为 val。是为下一个比其大的数。如果栈空，就把 -1 作为 val

- 把当前数入栈

- 循环结束后，用 nums1 中的数去 map 里查询生成结果

## 654, [最大二叉树](https://leetcode-cn.com/problems/maximum-binary-tree/)

- 递归 + 分治。递归终止条件：数组为空或者只有一个元素。

- 使用 max 函数返回数组中最大元素的 index
