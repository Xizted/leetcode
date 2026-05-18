class ListNode {
     val: number
     next: ListNode | null
     constructor(val?: number, next?: ListNode | null) {
         this.val = (val===undefined ? 0 : val)
         this.next = (next===undefined ? null : next)
     }
}


function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {

    let dummyNode = new ListNode(0)
    let carry = 0
    let current = dummyNode

    while (l1 != null || l2 != null) {
        let a = l1?.val ?? 0
        let b = l2?.val ?? 0

        if (l1) {
            l1 = l1.next
        }

        if (l2) {
            l2 = l2.next
        }

        let sum = a + b + carry
        let result = sum % 10
        carry = Math.trunc(sum / 10)

        current.next = new ListNode(result)
        current = current.next
    }

    if (carry > 0) {
        current.next = new ListNode(carry)
    }

    return dummyNode.next
    
};