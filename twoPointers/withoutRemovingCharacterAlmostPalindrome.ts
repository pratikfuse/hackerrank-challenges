
interface IsPalindromeResult {
    isStringPalindrome: boolean;
    wrongIndex: {
        left: number;
        right: number;
    };
}
const isPalindrome = (str: string, leftPointer: number, rightPointer: number): IsPalindromeResult => {
    // move inwards palindrome check

    while (leftPointer < rightPointer) {
        console.log(leftPointer, rightPointer);
        if (str[leftPointer] !== str[rightPointer]) {
            console.log(leftPointer, rightPointer, str[leftPointer], str[rightPointer])
            return {
                isStringPalindrome: false,
                wrongIndex: {
                    left: leftPointer,
                    right: rightPointer
                }
            };
        }
        leftPointer++;
        rightPointer--;
    }
    return {
        isStringPalindrome: true,
        wrongIndex: {
            left: leftPointer,
            right: rightPointer
        }
    }
}



const almostPalindrome = (str: string): boolean => {
    const { isStringPalindrome, wrongIndex } = isPalindrome(str, 0, str.length - 1);
    if (isStringPalindrome) {
        return true;
    }

    return isPalindrome(str, wrongIndex.left + 1, wrongIndex.right).isStringPalindrome || isPalindrome(str, wrongIndex.left, wrongIndex.right - 1).isStringPalindrome;
    // if (leftRemovalCheck.isStringPalindrome) {
    //     return true;
    // }
    // const rightRemovalCheck = isPalindrome(str, wrongIndex.left, wrongIndex.right - 1);
    // if (rightRemovalCheck.isStringPalindrome) {
    //     return true;
    // }

    // return false;
}

const s = almostPalindrome("raceacar");
console.log(s);