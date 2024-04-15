var canJump = function (nums) {
  const target = nums.length - 1;
  let reachable = 0;

  for (let index = 0; index < nums.length && index <= reachable; index++) {
    reachable = Math.max(reachable, index + nums[index]);
    if (reachable >= target) return true;
  }

  return false;
};
