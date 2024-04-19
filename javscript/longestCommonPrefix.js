/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
    let left = 0
    let right = 1

    let commonPrefix = new Map()
    
    for (let s of strs) {
        // get common prefix
        // if prefix is there then, move right to next, set the prefix var
        // no prefix then move left and right to next
    }
};

const case1 = ["flower","flow","flight"]
longestCommonPrefix(case1)