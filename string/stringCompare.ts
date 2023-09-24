

const getDerivedString = (string: string): string => {
    let updatedString: string[] = [];

    for (let char of string.split('')) {

        if (char === "#") {
            updatedString.pop();
            continue;
        }
        updatedString.push(char);
    }
    return updatedString.join("");
}

const compareStrings = (string1: string, string2: string): boolean => {

    const derivedString1 = getDerivedString(string1);
    const derivedString2 = getDerivedString(string2);
    console.log(derivedString1 === derivedString2);
    return false;
}

const optimalSolution = (string1: string, string2: string): boolean => {
    let p1 = string1.length - 1, p2 = string2.length - 1;
    while (p1 >= 0 || p2 >= 0) {
        if (string1[p1] === "#" || string2[p2] === "#") {
            if (string1[p1] === "#") {
                let backCount = 2;
                while (backCount > 0) {
                    p1--;
                    backCount--;
                    if (string1[p1] === "#") backCount += 2;
                }
            }

            if (string2[p2] === "#") {
                let backCount = 2;
                while (backCount > 0) {
                    p2--;
                    backCount--;
                    if (string2[p2] === "#") backCount += 2;
                }
            }
        } else {
            if (string1[p1] !== string2[p2]) {
                return false;
            }
            p1--;
            p2--;
        }
    }
    return true;
}


const isEqual = optimalSolution("ab#c#d", "ab#c#de#");
console.log(isEqual);