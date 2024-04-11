/**
 * Essentially we just loop through each item, if we find
 * a match, move it to the end because the response wants
 * all matches at the end. We actually swap the last item
 * with the current one that gets replaced. Then we
 * reduce the window size because we "removed" an item.
 *
 * This also keeps the Big O at:
 * Time: O(n) because we iterate through the entire array once.
 * Space: O(1) because we are not using any extra space that grows with the input size.
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function (nums, val) {
  var numsLen = nums.length;
  var swappedCount = 0;
  for (var i = 0; i < numsLen; ) {
    // did we find a match? swap the last item with this _
    // also reduce our counter length since we "removed" one
    // don't increment i because we have to re-check the one we swapped
    if (nums[i] == val) {
      nums[i] = nums[numsLen - 1]; // swapped
      nums[numsLen - 1] = "_"; // "removed"
      numsLen--;
      swappedCount++;
    } else {
      i++;
    }
  }
  return nums.length - swappedCount;
};
