var MinStack = function () {
  this.list = [];
};

/**
 * @param {number} val
 * @return {void}
 */
MinStack.prototype.push = function (val) {
  this.list.push(val);
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function () {
  this.list.pop();
};

/**
 * @return {number}
 */
MinStack.prototype.top = function () {
  return this.list[this.list.length - 1];
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function () {
  return Math.min(...this.list);
};
