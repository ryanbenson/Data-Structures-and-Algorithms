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

/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWordLoop = function (s) {
  s = s.trim();
  if (s.length === 0) {
    return 0;
  }
  if (s.length === 1) {
    return 1;
  }
  var spaceI = 0;
  for (var i = s.length - 1; i >= 0; i--) {
    if (s[i] === " ") {
      spaceI = i;
      break;
    }
  }
  if (spaceI === 0) {
    return s.length;
  }
  return s.length - 1 - spaceI;
};
