/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  var profit = 0;
  var pricesLen = prices.length;
  for (var i = 1; i < pricesLen; i++) {
    var curPrice = prices[i];
    var lastPrice = prices[i - 1];
    if (curPrice > lastPrice) {
      profit += curPrice - lastPrice;
    }
  }
  return profit;
};
