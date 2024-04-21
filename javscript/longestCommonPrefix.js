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

/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefixV2 = function(strs) {
    let result = '',
        prefix = strs[0];
    
    for (let i = 0; i < prefix.length; i++) {
      if (strs.some(str => str[i] !== prefix[i])) {
        break;
      }
      result += prefix[i];
    }
    
    console.log(result)
    return result
  };

const case1 = ["flower","flow","flight"]
longestCommonPrefixV2(case1)