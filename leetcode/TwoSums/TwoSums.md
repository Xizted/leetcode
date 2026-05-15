---
tags:
  - Resolved
  - Easy
---
### Hint

Given an array of integers `nums` and an integer `target`, return _indices of the two numbers such that they add up to `target`_.

You may assume that each input would have **_exactly_ one solution**, and you may not use the _same_ element twice.

You can return the answer in any order.

### **Example 1:**


```
**Input:** nums = [2,7,11,15], target = 9
**Output:** [0,1]
**Explanation:** Because nums[0] + nums[1] == 9, we return [0, 1].

**Example 2:**

**Input:** nums = [3,2,4], target = 6
**Output:** [1,2]

**Example 3:**

**Input:** nums = [3,3], target = 6
**Output:** [0,1]

**Constraints:**

- `2 <= nums.length <= 104`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
- **Only one valid answer exists.**
```

**Follow-up:** Can you come up with an algorithm that is less than `O(n2)` time complexity?

### Solución

1. Creamos un map para guardar los números recorridos: value => index
2. Recorremos el arreglo de número y obtenemos el valor buscado para obtener el target: target - currentValue
3. Validamos si el número existe en el map:
	- Si existe, devolvemos en un arreglo el indicé guardado y el indicé actual
	- Caso contrario, guardaremos en el map el valor actual como key y el indicé como valor.

### Go

```go
func twoSum(nums []int, target int) []int {
	
	if empty := len(nums) == 0; empty {
		return nil
	}
	
	numsSaved := make(map[int]int, len(nums))
	
	for currentIndex, currentValue := range nums {
		valueSearched := target - currentValue
		if value, exist := numsSaved[valueSearched]; exist {
			return []int{value, currentIndex}
		}
		
		numsSaved[currentValue] = currentIndex
	}
	
	return nil
}
```

### Javascript/Typescript

```ts
function twoSum(nums: number[], target: number): number[] {
	if (nums.length === 0) return []
	
	const numsSaved = new Map<number,number>()
	
	for (let currentIndex = 0; currentIndex < nums.length; i++) {
		const valueSearched = target - nums[currentIndex]
		if (numsSaved.has(valueSearched)) {
			return [numsSaved.get(valueSearched), currentIndex]
		}
		numsSaved.set(nums[i], currentIndex)
	}
	
	return []
}
```