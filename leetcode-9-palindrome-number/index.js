/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
  x = x.toString();
  const xLen = x.length;
  if (xLen === 1) {
    return true;
  }
  if (xLen === 2) {
    return x[0] === x[1];
  }
  const stoppingPointer = Math.floor(x.length / 2);
  for (let i = 0; i < stoppingPointer; i++) {
    if (x[i] !== x[xLen - (1 + i)]) {
      return false;
    }
  }
  return true;
};
