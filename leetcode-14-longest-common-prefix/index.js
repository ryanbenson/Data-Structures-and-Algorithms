/**
 * Idea is simple. Just grab the first letter, if one is found
 * then check each item for it
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
  let longestPrefix = "";
  const totalStrings = strs.length;
  let i = 0;
  while (true) {
    let letter = strs[0]?.[i];
    if (!letter) {
      break;
    }
    let isMatch = true;
    for (let k = 0; k < totalStrings; k++) {
      if (strs[k]?.[i] !== letter) {
        isMatch = false;
        break;
      }
    }
    if (isMatch) {
      longestPrefix += letter;
    } else {
      break;
    }
    i++;
  }
  return longestPrefix;
};
