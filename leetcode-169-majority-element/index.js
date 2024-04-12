/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
  const numsLen = nums.length;
  if (numsLen === 1) {
    return nums[0];
  }
  nums.sort(function (a, b) {
    return a - b;
  });
  let curNum = nums[0];
  let curNumberCount = 1;
  let curMajorityNum = nums[0];
  let curMajorityCount = 1;
  for (let i = 1; i < numsLen; i++) {
    if (nums[i] === curNum) {
      curNumberCount++;
    } else {
      curNum = nums[i];
      curNumberCount = 1;
    }
    if (curNum === curMajorityNum) {
      curMajorityCount++;
    }
    if (curNumberCount > curMajorityCount) {
      curMajorityNum = curNum;
    }
  }
  return curMajorityNum;
};
