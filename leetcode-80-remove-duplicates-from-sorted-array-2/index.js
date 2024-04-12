/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  const numsLen = nums.length;
  // don't bother if we have too small of an array
  if (numsLen <= 1) {
    return numsLen;
  }
  var i = 0;
  var j = 0;
  // if the first and second numbers are the same,
  // don't bother using the first item in the loop, we can skip it
  if (nums[0] == nums[1]) {
    i = 1;
    j = 1;
  }
  for (j; j < numsLen; j++) {
    if (nums[j] != nums[i]) {
      i += 1;
      nums[i] = nums[j];
      if (nums[j] == nums[j + 1]) {
        nums[i + 1] = nums[j];
        i++;
      }
    }
  }
  return i + 1;
};
