/**
 * This one makes a lot more legible sense, but it's O((m+n)log(m+n))
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
  // get the chunks we need from nums2
  var nums2Els = nums2.slice(0, n);
  // chop off what we don't need from nums1
  nums1.splice(m, nums1.length);
  // merge the two together
  nums1.splice(nums1.length, 0, ...nums2Els);
  // sort numbers properly
  nums1.sort(function (a, b) {
    return a - b;
  });
};

/**
 * This one is O(m+n)
 * It's harder to read and understand, but it works.
 * Since both arrays provided are already sorted, we can start at the end
 * You keep adding in the numbers, to fill in the spaces, then you end
 * up start moving numbers until you are finished with both arrays
 * then they're sorted. In action:
 * nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
 *
 * [ 1, 2, 3, 0, 0, 6 ]
 * [ 1, 2, 3, 0, 5, 6 ]
 * [ 1, 2, 3, 3, 5, 6 ]
 * [ 1, 2, 2, 3, 5, 6 ]
 * [ 1, 2, 2, 3, 5, 6 ]
 * [ 1, 2, 2, 3, 5, 6 ]
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge2 = function (nums1, m, nums2, n) {
  let i = m - 1;
  let j = n - 1;
  let k = nums1.length - 1;

  while (k >= 0) {
    if (nums1[i] > nums2[j]) {
      if (nums1[i] !== undefined) {
        nums1[k] = nums1[i--];
      }
    } else {
      if (nums2[j] !== undefined) {
        nums1[k] = nums2[j--];
      }
    }
    k--;
  }
  console.log(nums1);
};
