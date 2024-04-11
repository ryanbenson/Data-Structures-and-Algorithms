/**
 * See the readme, but basically:
 * You have a pointer for the items in the array to replace,
 * and you only increment that pointer whenever
 * a new number comes in with the main loop.
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  const numsLen = nums.length;
  // don't bother if we have too small of an array
  if (numsLen <= 1) {
    return numsLen;
  }
  var lastNumberIndex = 1;
  for (var i = 1; i < numsLen; i++) {
    if (nums[i - 1] != nums[i]) {
      nums[lastNumberIndex] = nums[i];
      lastNumberIndex++;
    }
  }
  return lastNumberIndex;
};
