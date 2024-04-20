/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function (s) {
  if (s.length === 0) {
    return 0;
  }
  if (s.length === 1) {
    return 1;
  }
  s = s.trimEnd();
  var words = s.split(" ");
  return words[words.length - 1].length;
};
