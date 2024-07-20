// 迭代
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    for cur := head; cur != nil; {
        next := cur.Next
        cur.Next = pre
        pre, cur = cur, next
    }
    return pre
}

// // 递归
// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }
//     p := head.Next
//     newHead := reverseList(head.Next)
//     p.Next, head.Next = head, nil
//     return newHead
// }
