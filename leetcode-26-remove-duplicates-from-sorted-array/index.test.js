const test = require("ava");
const { removeDuplicates } = require("./index");

test("#removeDuplicates when given a small sample", (t) => {
  const sample = [1, 2, 2];
  const valid = [1, 2];
  const expectedNums = 2;
  const result = removeDuplicates(sample);
  t.is(result, expectedNums);
  for (var i = 0; i < expectedNums; i++) {
    t.is(sample[i], valid[i]);
  }
});

test("#removeDuplicates when given a bigger sample", (t) => {
  const sample = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
  const valid = [0, 1, 2, 3, 4, 2, 2, 3, 3, 4];
  const expectedNums = 5;
  const result = removeDuplicates(sample);
  t.is(result, expectedNums);
  for (var i = 0; i < expectedNums; i++) {
    t.is(sample[i], valid[i]);
  }
});

test("#removeDuplicates when given a too small of a sample", (t) => {
  const sample = [1];
  const valid = [1];
  const expectedNums = 1;
  const result = removeDuplicates(sample);
  t.is(result, expectedNums);
  for (var i = 0; i < expectedNums; i++) {
    t.is(sample[i], valid[i]);
  }
});
