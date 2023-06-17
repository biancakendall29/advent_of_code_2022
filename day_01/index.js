const fs = require("fs");

const textIn = fs.readFileSync("./input.txt", "utf8");
const arr = textIn.split("\n");

let sum = 0;
let totals = [];
for (let i = 0; i < arr.length; i++) {
  sum += Number(arr[i]);
  if (arr[i] === "") {
    totals.push(sum);
    sum = 0;
    continue;
  }
}

const topThree = totals.sort((a, b) => b - a).slice(0, 3);

// Part 1
console.log("Maximum: " + topThree[0]);

// Part 2
console.log(
  "Sum of top 3: " + topThree.reduce((partialSum, a) => partialSum + a, 0)
);
