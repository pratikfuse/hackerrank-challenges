
interface IsPalindromeResult {
    isStringPalindrome: boolean;
    wrongIndex: {
        left: number;
        right: number;
    };
}
const isPalindrome = (str: string): IsPalindromeResult => {
    // move inwards palindrome check

    let leftPointer = 0;
    let rightPointer = str.length - 1;

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

function removeCharacter(str: string, position: number): string {
    const newString = str.split("").filter((_, i) => i !== position).join("");
    return newString
}

const almostPalindrome = (str: string): boolean => {
    const { isStringPalindrome, wrongIndex } = isPalindrome(str);
    if (isStringPalindrome) {
        return true;
    }

    const leftRemovalCheck = isPalindrome(removeCharacter(str, wrongIndex.left));
    if (leftRemovalCheck.isStringPalindrome) {
        return true;
    }
    const rightRemovalCheck = isPalindrome(removeCharacter(str, wrongIndex.right));
    if (rightRemovalCheck.isStringPalindrome) {
        return true;
    }

    return false;
}

const s = almostPalindrome("raceacar");
console.log(s);



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

function removeCharacter(str: string, position: number): string {
    const newString = str.split("").filter((_, i) => i !== position).join("");
    return newString
}

const almostPalindrome = (str: string): boolean => {
    const { isStringPalindrome, wrongIndex } = isPalindrome(str, 0, str.length - 1);
    if (isStringPalindrome) {
        return true;
    }

    const leftRemovalCheck = isPalindrome(removeCharacter(str, wrongIndex.left), wrongIndex.left + 1, wrongIndex.right);
    if (leftRemovalCheck.isStringPalindrome) {
        return true;
    }
    const rightRemovalCheck = isPalindrome(removeCharacter(str, wrongIndex.right), wrongIndex.left + 1, wrongIndex.right);
    if (rightRemovalCheck.isStringPalindrome) {
        return true;
    }

    return false;
}

const s = almostPalindrome("raceacar");
console.log(s);