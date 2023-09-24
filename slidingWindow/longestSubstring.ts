


function longestSubstring(str: string) {

    if (str.length <= 1) {
        return str.length;
    }

    let leftPointer = 0;
    let longestSubstringCount = 0;

    const seenCharacters: { [key: string]: number } = {};

    for (let rightPointer = 0; rightPointer < str.length; rightPointer++) {
        const currentCharacter = str[rightPointer];
        const seenCharacterIndex = seenCharacters[currentCharacter];

        // character found in hash map

        if (seenCharacterIndex >= leftPointer) {
            leftPointer = seenCharacterIndex + 1;
        }

        seenCharacters[currentCharacter] = rightPointer;

        longestSubstringCount = Math.max(longestSubstringCount, rightPointer - leftPointer + 1)


    }

    return longestSubstringCount;
}

longestSubstring("abcbdabc");