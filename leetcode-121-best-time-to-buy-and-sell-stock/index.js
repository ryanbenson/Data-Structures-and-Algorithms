/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  const pricesLen = prices.length;
  // don't bother looping if there's no reason to
  if (!pricesLen || pricesLen == 1) {
    return 0;
  }
  if (pricesLen == 2) {
    const margin = prices[1] - prices[0];
    if (margin < 0) {
      return 0;
    }
  }
  let highestMargin = 0;
  // don't bother checking the last item
  for (let i = 0; i < pricesLen - 1; i++) {
    let margin = Math.max(...prices.slice(i + 1)) - prices[i];
    if (margin > highestMargin) {
      highestMargin = margin;
    }
  }

  return highestMargin;
};

/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfitEfficient = function (prices) {
  const pricesLen = prices.length;
  // don't bother looping if there's no reason to
  if (!pricesLen || pricesLen == 1) {
    return 0;
  }
  if (pricesLen == 2) {
    const margin = prices[1] - prices[0];
    if (margin < 0) {
      return 0;
    }
  }
  let highestMargin = 0;
  let min = Number.MAX_SAFE_INTEGER;
  // don't bother checking the last item
  for (let i = 0; i < pricesLen; i++) {
    if (prices[i] < min) {
      min = prices[i];
    } else {
      const margin = prices[i] - min;
      if (margin > highestMargin) {
        highestMargin = margin;
      }
    }
  }

  return highestMargin;
};
