function isAnagram(s: string, t: string): boolean {

    if (s.length !== t.length) return false

    const charInventory = new Map<string, number>()

    for(let i = 0; i < s.length; i++) {
        const currentChar: string = s[i];
        if (!charInventory.has(currentChar)) {
            charInventory.set(currentChar, 1)
            continue
        }
        charInventory.set(currentChar, charInventory.get(currentChar)! + 1)
    }

    for(let i = 0; i < t.length; i++) {
        const currentChar: string = t[i]

        if (!charInventory.has(currentChar)) {
            return false
        }

        if (charInventory.get(currentChar) === 0) {
            return false
        }

        charInventory.set(currentChar, charInventory.get(currentChar)! - 1)
    }
    
    return true
};