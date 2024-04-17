/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    let k = 0
    let numsv2 = []
    for (let i = 0; i < nums.length; i ++){
        if(numsv2.includes(nums[i])) {
            k++
            continue
        }
        numsv2.push(nums[i])
    }

    return k
};

const removeDuplicatesV2 = (nums) => {
    console.log(nums.length - [...new Set(nums)].length)
    console.log([...new Set(nums)])
}

const case1 = [0,0,1,1,1,2,2,3,3,4]
removeDuplicates(case1)
removeDuplicatesV2(case1)