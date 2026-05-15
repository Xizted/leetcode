function twoSum(nums: number[], target: number): number[] {
	if (nums.length === 0) return []
	
	const numsSaved = new Map<number,number>()
	
	for (let currentIndex = 0; currentIndex < nums.length; currentIndex++) {
		const valueSearched = target - nums[currentIndex]
		if (numsSaved.has(valueSearched)) {
			return [numsSaved.get(valueSearched)!, currentIndex]
		}
		numsSaved.set(nums[currentIndex], currentIndex)
	}
	
	return []
}

